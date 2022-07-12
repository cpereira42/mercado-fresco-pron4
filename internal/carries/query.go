package carries

const (
	CreateCarry = `INSERT INTO carries (cid, company_name, address, telephone, locality_id) VALUES (?, ?, ?, ?, ?)`

	GetByIDReport = `SELECT locality_id, locality_name, COUNT(carries.locality_id)
	FROM localities
	INNER JOIN carries ON localities.id = carries.locality_id
	WHERE localities.id = ?`

	GetAllReport = `SELECT locality_id, locality_name, COUNT(carries.locality_id)
	FROM localities
	INNER JOIN carries ON localities.id = carries.locality_id
	GROUP BY localities.id`
)
