package repository

import (
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/rdbms"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

const (
	allocationClassesFromAllSourcesSQL = `
		WITH class_hierarchy_position AS (
			SELECT
				p.id AS portfolio_id,
				(
					SELECT pos - 1
					FROM jsonb_array_elements(p.allocation_structure->'hierarchy') WITH ORDINALITY AS t(elem, pos)
					WHERE elem->>'field' = 'class'
					LIMIT 1
				) AS class_position
			FROM portfolio p
			WHERE p.id = {:portfolioId}
		)
		SELECT DISTINCT class
		FROM (
			SELECT pa.class
			FROM portfolio_allocation_fact pa
			WHERE pa.portfolio_id = {:portfolioId}
			UNION
			SELECT pa.hierarchical_id[chp.class_position + 1] AS class
			FROM planned_allocation pa
			JOIN allocation_plan ap ON pa.allocation_plan_id = ap.id
			CROSS JOIN class_hierarchy_position chp
			WHERE ap.portfolio_id = {:portfolioId}
				AND chp.class_position IS NOT NULL
				AND pa.hierarchical_id[chp.class_position + 1] IS NOT NULL
		) AS combined_classes
		ORDER BY class ASC
	`
)

type AllocationRDBMSRepository struct {
	dbAdapter rdbms.RepositoryRDBMSAdapter
}

// FindAvailableAllocationClassesFromAllSources retrieves unique allocation classes
// from both portfolio_allocation_fact and planned_allocation tables. For planned_allocation,
// it extracts the class value from the hierarchical_id array using the position defined in
// the portfolio's allocation_structure.
//
// Authored by: GitHub Copilot
func (repository *AllocationRDBMSRepository) FindAvailableAllocationClassesFromAllSources(
	portfolioId int64,
) ([]string, error) {

	rows, err := repository.dbAdapter.BuildQuery(allocationClassesFromAllSourcesSQL).
		AddParam("portfolioId", portfolioId).
		Build().GetRows()

	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error querying allocation classes from all sources",
			repository,
		)
	}

	return repository.scanAllocationClassesRows(rows)
}

// scanAllocationClassesRows scans the result rows to extract allocation classes.
//
// Authored by: GitHub Copilot
func (repository *AllocationRDBMSRepository) scanAllocationClassesRows(
	rows *dbx.Rows,
) ([]string, error) {

	var queryResult = make([]string, 0)
	for rows.Next() {

		var class string
		err := rows.Scan(&class)
		if err != nil {
			return nil, infra.PropagateAsAppErrorWithNewMessage(
				err,
				"Error scanning allocation class",
				repository,
			)
		}

		queryResult = append(queryResult, class)
	}
	return queryResult, nil
}

// BuildAllocationRepository creates a new AllocationRDBMSRepository instance.
//
// Authored by: GitHub Copilot
func BuildAllocationRepository(dbAdapter rdbms.RepositoryRDBMSAdapter) *AllocationRDBMSRepository {
	return &AllocationRDBMSRepository{dbAdapter: dbAdapter}
}
