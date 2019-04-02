// VulcanizeDB
// Copyright © 2019 Vulcanize

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

package contenthash_changed

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	repo "github.com/vulcanize/vulcanizedb/libraries/shared/repository"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"

	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
)

type ContenthashChangedRepository struct {
	db *postgres.DB
}

func (repository *ContenthashChangedRepository) SetDB(db *postgres.DB) {
	repository.db = db
}

func (repository ContenthashChangedRepository) Create(headerID int64, models []interface{}) error {
	tx, dBaseErr := repository.db.Beginx()
	if dBaseErr != nil {
		return dBaseErr
	}
	for _, model := range models {
		contentModel, ok := model.(ContenthashChangedModel)
		if !ok {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Error("failed to rollback ", rollbackErr)
			}
			return fmt.Errorf("model of type %T, not %T", model, ContenthashChangedModel{})
		}

		_, execErr := tx.Exec(
			`INSERT into ens.contenthash_changed (header_id, resolver, node, hash, log_idx, tx_idx, raw_log)
        			VALUES($1, $2, $3, $4, $5, $6, $7)
					ON CONFLICT (header_id, tx_idx, log_idx) DO UPDATE SET resolver = $2, node = $3, hash = $4, raw_log = $7;`,
			headerID, contentModel.Resolver, contentModel.Node, contentModel.Hash, contentModel.LogIndex, contentModel.TransactionIndex, contentModel.Raw,
		)
		if execErr != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Error("failed to rollback ", rollbackErr)
			}
			return execErr
		}
	}

	checkHeaderErr := repo.MarkHeaderCheckedInTransaction(headerID, tx, constants.ContenthashChangedChecked)
	if checkHeaderErr != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Error("failed to rollback ", rollbackErr)
		}
		return checkHeaderErr
	}

	return tx.Commit()
}

func (repository ContenthashChangedRepository) MarkHeaderChecked(headerID int64) error {
	return repo.MarkHeaderChecked(headerID, repository.db, constants.ContenthashChangedChecked)
}

func (repository ContenthashChangedRepository) MissingHeaders(startingBlockNumber int64, endingBlockNumber int64) ([]core.Header, error) {
	return repo.MissingHeaders(startingBlockNumber, endingBlockNumber, repository.db, constants.ContenthashChangedChecked)
}

func (repository ContenthashChangedRepository) RecheckHeaders(startingBlockNumber int64, endingBlockNumber int64) ([]core.Header, error) {
	return repo.RecheckHeaders(startingBlockNumber, endingBlockNumber, repository.db, constants.ContenthashChangedChecked)
}
