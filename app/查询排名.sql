SELECT
	a.NAME,
	a.vote_count,
CASE
		
		WHEN @rowtotal = a.vote_count THEN
		@rownum 
		WHEN @rowtotal := a.vote_count THEN
		@rownum := @rownum + 1 
		WHEN @rowtotal = 0 THEN
		@rownum := @rownum + 1 
	END AS rownum 
FROM
	( SELECT NAME, vote_count FROM construction_development_h5_vote_data ORDER BY vote_count DESC ) AS a,
	( SELECT @rownum := 0, @rowtotal := NULL ) r


    