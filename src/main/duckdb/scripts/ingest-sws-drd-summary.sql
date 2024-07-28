CREATE TEMP TABLE sws_summary AS
    SELECT *
    FROM read_csv(
        '/duckdb/input/us-complete-portfolio-summary.csv'--,
--         columns = {
--             'asset': 'text',
--             'total_bought': ''
--         }
    );
    
SELECT * FROM sws_summary;