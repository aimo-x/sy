SELECT * FROM ( SELECT id,( @rowNum := @rowNum + 1 ) AS rowNo FROM construction_development_h5_vote_data,( SELECT ( @rowNum := 0 ) ) b 
	ORDER BY
		vote_count DESC 
	) c 
WHERE
	id = 30;