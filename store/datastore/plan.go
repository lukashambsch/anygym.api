package datastore

import (
	"fmt"

	"github.com/lukashambsch/gym-all-over/models"
	"github.com/lukashambsch/gym-all-over/store"
)

func GetPlanList(where string) ([]models.Plan, error) {
	var (
		plans []models.Plan
		plan  models.Plan
	)

	query := fmt.Sprintf("%s %s", getPlanListQuery, where)
	rows, err := store.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&plan.PlanId, &plan.PlanName, &plan.Price)
		plans = append(plans, plan)
		if err != nil {
			return nil, err
		}
	}
	defer rows.Close()

	return plans, nil
}

func GetPlanCount(where string) (*int, error) {
	var count int

	query := fmt.Sprintf("%s %s", getPlanCountQuery, where)
	row := store.DB.QueryRow(query)
	err := row.Scan(&count)
	if err != nil {
		return nil, err
	}

	return &count, nil
}

func GetPlan(planId int64) (*models.Plan, error) {
	var plan models.Plan

	row := store.DB.QueryRow(getPlanQuery, planId)
	err := row.Scan(&plan.PlanId, &plan.PlanName, &plan.Price)
	if err != nil {
		return nil, err
	}

	return &plan, nil
}

func CreatePlan(plan models.Plan) (*models.Plan, error) {
	var created models.Plan

	row := store.DB.QueryRow(createPlanQuery, plan.PlanName, plan.Price)
	err := row.Scan(&created.PlanId, &created.PlanName, &created.Price)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func UpdatePlan(planId int64, plan models.Plan) (*models.Plan, error) {
	var updated models.Plan

	row := store.DB.QueryRow(updatePlanQuery, plan.PlanName, plan.Price, planId)
	err := row.Scan(&updated.PlanId, &updated.PlanName, &updated.Price)
	if err != nil {
		return nil, err
	}

	return &updated, nil
}

func DeletePlan(planId int64) error {
	stmt, err := store.DB.Prepare(deletePlanQuery)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(planId)
	if err != nil {
		return err
	}

	return nil
}

const getPlanListQuery = `
SELECT *
FROM plans
`

const getPlanQuery = `
SELECT *
FROM plans
WHERE plan_id = $1
`

const createPlanQuery = `
INSERT INTO plans (plan_name, price)
VALUES ($1, $2)
RETURNING plan_id, plan_name, price
`

const updatePlanQuery = `
UPDATE plans
SET plan_name = $1, price = $2
WHERE plan_id = $3
RETURNING plan_id, plan_name, price
`

const deletePlanQuery = `
DELETE
FROM plans
WHERE plan_id = $1
`

const getPlanCountQuery = `
SELECT count(*)
FROM plans
`
