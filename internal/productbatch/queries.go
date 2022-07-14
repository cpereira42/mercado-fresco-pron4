package productbatch

const (
	SqlCreatePB          = `INSERT INTO mercadofresco.products_batches (batch_number,current_quantity,current_temperature,due_date,initial_quantity,manufacturing_date,manufacturing_hour,minimum_temperature,product_id,section_id) VALUES (?,?,?,?,?,?,?,?,?,?)`
	SqlBatcheNumber      = `SELECT batch_number FROM mercadofresco.products_batches WHERE batch_number=?`
	SqlrelatorioTodo     = `SELECT s.id AS 'section_id', s.section_number AS 'section_number', count(s.id) AS 'products_count' FROM bgow1s413.products_batches as pbs INNER JOIN bgow1s413.sections AS s ON s.id = pbs.section_id INNER JOIN bgow1s413.products AS pcts ON pbs.product_id = pcts.id group by s.id having count(s.id) >= 0;`
	SqlrelatorioSectioId = `SELECT s.id AS 'section_id', s.section_number AS 'section_number', count(s.id) AS 'products_count' FROM bgow1s413.products_batches as pbs INNER JOIN bgow1s413.sections AS s ON s.id = pbs.section_id INNER JOIN bgow1s413.products AS pcts ON pbs.product_id = pcts.id where s.id=? group by s.id having count(s.id) >= 0`
)
