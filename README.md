# ws-load-test

--host 172.31.25.37:8080 --protocol ws --requestCount 1000 --gapTime 500 --path /mds --lifeTime 1 --strategy ping_pong --writeTime 0.5

--host 172.31.25.37:8080 --protocol ws --requestCount 1000 --gapTime 500 --path /smart-stream --lifeTime 10 --strategy exchange_tick --writeTime 9999

--host 172.31.25.37:8080 --protocol ws --requestCount 1000 --gapTime 500 --path /mds --lifeTime 1 --messageText all_ticks --writeTime 500   

