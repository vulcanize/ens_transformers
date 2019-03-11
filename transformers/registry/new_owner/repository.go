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

package new_owner

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	repo "github.com/vulcanize/vulcanizedb/libraries/shared/repository"
	"github.com/vulcanize/vulcanizedb/pkg/core"
	"github.com/vulcanize/vulcanizedb/pkg/datastore/postgres"

	"github.com/vulcanize/ens_transformers/transformers/shared/constants"
)

type NewOwnerRepository struct {
	db *postgres.DB
}

func (repository *NewOwnerRepository) SetDB(db *postgres.DB) {
	repository.db = db
}

func (repository NewOwnerRepository) Create(headerID int64, models []interface{}) error {
	tx, dBaseErr := repository.db.Begin()
	if dBaseErr != nil {
		return dBaseErr
	}
	for _, model := range models {
		ownerModel, ok := model.(NewOwnerModel)
		if !ok {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Error("failed to rollback ", rollbackErr)
			}
			return fmt.Errorf("model of type %T, not %T", model, NewOwnerModel{})
		}

		_, execErr := tx.Exec(
			`INSERT into ens.new_owner (header_id, node, label, owner, subnode, log_idx, tx_idx, raw_log)
        			VALUES($1, $2, $3, $4, $5, $6, $7, $8)
					ON CONFLICT (header_id, tx_idx, log_idx) DO UPDATE SET node = $2, label = $3, owner = $4, subnode = $5, raw_log = $8;`,
			headerID, ownerModel.Node, ownerModel.Label, ownerModel.Owner, ownerModel.Subnode, ownerModel.LogIndex, ownerModel.TransactionIndex, ownerModel.Raw,
		)
		if execErr != nil {
			rollbackErr := tx.Rollback()
			if rollbackErr != nil {
				log.Error("failed to rollback ", rollbackErr)
			}
			return execErr
		}
	}

	checkHeaderErr := repo.MarkHeaderCheckedInTransaction(headerID, tx, constants.NewOwnerChecked)
	if checkHeaderErr != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			log.Error("failed to rollback ", rollbackErr)
		}
		return checkHeaderErr
	}

	return tx.Commit()
}

func (repository NewOwnerRepository) MarkHeaderChecked(headerID int64) error {
	return repo.MarkHeaderChecked(headerID, repository.db, constants.NewOwnerChecked)
}

func (repository NewOwnerRepository) MissingHeaders(startingBlockNumber int64, endingBlockNumber int64) ([]core.Header, error) {
	return repo.MissingHeaders(startingBlockNumber, endingBlockNumber, repository.db, constants.NewOwnerChecked)
}

func (repository NewOwnerRepository) RecheckHeaders(startingBlockNumber int64, endingBlockNumber int64) ([]core.Header, error) {
	return repo.RecheckHeaders(startingBlockNumber, endingBlockNumber, repository.db, constants.NewOwnerChecked)
}
