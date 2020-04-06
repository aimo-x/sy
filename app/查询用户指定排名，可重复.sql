SELECT
	* 
FROM
	(
	SELECT
		obj.id,
		obj.time_count,
	CASE
			
			WHEN @rowtotal = obj.time_count THEN
			@rownum 
			WHEN @rowtotal := obj.time_count THEN
			@rownum := @rownum + 1 
			WHEN @rowtotal = 0 THEN
			@rownum := @rownum + 1 
		END AS rownum 
	FROM
		( SELECT id, time_count FROM game_flop_logs ORDER BY time_count ) AS obj,
		( SELECT @rownum := 0, @rowtotal := NULL ) r 
	) c
WHERE
	id = 1370;