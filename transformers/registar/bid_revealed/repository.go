// VulcanizeDB
// Copyright © 2018 Vulcanize

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package bid_revealed

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	repo "github.com/vulcanize/vulcanizedb/libraries/shared/repository"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"

	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
)

type BidRevealedRepository struct {
	db *postgres.DB
}

func (repository *BidRevealedRepository) SetDB(db *postgres.DB) {
	repository.db = db
}

func (repository BidRevealedRepository) Create(headerID int64, models []interface{}) error {
	tx, dBaseErr := repository.db.Beginx()
	if dBaseErr != nil {
		return dBaseErr
	}
	for _, model := range models {
		bidModel, ok := model.(BidRevealedModel)
		if !ok {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Error("failed to rollback ", rollbackErr)
			}
			return fmt.Errorf("model of type %T, not %T", model, BidRevealedModel{})
		}

		_, execErr := tx.Exec(
			`INSERT into ens.bid_revealed (header_id, hash, owner, value, status, log_idx, tx_idx, raw_log)
        			VALUES($1, $2, $3, $4, $5, $6, $7, $8)
					ON CONFLICT (header_id, tx_idx, log_idx) DO UPDATE SET hash = $2, owner = $3, value = $4, status = $5, raw_log = $8;`,
			headerID, bidModel.Hash, bidModel.Owner, bidModel.Value, bidModel.Status, bidModel.LogIndex, bidModel.TransactionIndex, bidModel.Raw,
		)
		if execErr != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Error("failed to rollback ", rollbackErr)
			}
			return execErr
		}
	}

	checkHeaderErr := repo.MarkHeaderCheckedInTransaction(headerID, tx, constants.BidRevealedChecked)
	if checkHeaderErr != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Error("failed to rollback ", rollbackErr)
		}
		return checkHeaderErr
	}

	return tx.Commit()
}

func (repository BidRevealedRepository) MarkHeaderChecked(headerID int64) error {
	return repo.MarkHeaderChecked(headerID, repository.db, constants.BidRevealedChecked)
}

func (repository BidRevealedRepository) MissingHeaders(startingBlockNumber int64, endingBlockNumber int64) ([]core.Header, error) {
	return repo.MissingHeaders(startingBlockNumber, endingBlockNumber, repository.db, constants.BidRevealedChecked)
}

func (repository BidRevealedRepository) RecheckHeaders(startingBlockNumber int64, endingBlockNumber int64) ([]core.Header, error) {
	return repo.RecheckHeaders(startingBlockNumber, endingBlockNumber, repository.db, constants.BidRevealedChecked)
}
