ACCESS_TOKEN ?= eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiI1Mzk1MGEwNS0xOWRiLTExZWItYjg5Ni0wMjQyYWMxMjAwMDIiLCJleHAiOjE2MDQyNDIyNTcsImlhdCI6MTYwNDIzODY1NywiaXNzIjoiYXV0aF9zZXJ2aWNlIiwibmJmIjoxNjA0MjM4NjU3fQ.LqcmdGIOo9d65DqLVkMO01N0WT-7mYXDfNZ8DoDMoII

.PHONEY: launch_1 launch_10 launch_100 launch_1000

launch_1:
	docker run --rm --net=host \
		williamyeh/wrk  \
		-H "Authorization: ${ACCESS_TOKEN}" \
		-t 1 \
		-c 1 \
		-d 1m \
		http://localhost:9999/questionnaires?prefix=ma ;\

launch_10:
	docker run --rm --net=host \
		williamyeh/wrk  \
		-H "Authorization: ${ACCESS_TOKEN}" \
		-t 10 \
		-c 10 \
		-d 1m \
		http://localhost:9999/questionnaires?prefix=ma ;\

launch_100:
	docker run --rm --net=host \
		williamyeh/wrk  \
		-H "Authorization: ${ACCESS_TOKEN}" \
		-t 10 \
		-c 100 \
		-d 1m \
		http://localhost:9999/questionnaires?prefix=ma ;\

launch_1000:
	docker run --rm --net=host \
		williamyeh/wrk  \
		-H "Authorization: ${ACCESS_TOKEN}" \
		-t 10 \
		-c 1000 \
		-d 1m \
		http://localhost:9999/questionnaires?prefix=ma ;\
