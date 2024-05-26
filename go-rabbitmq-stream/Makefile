docker-rabbitmq:
	@docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.9-manegement

docker-rabbitmq-delete-events:
	@docker exec -it rabbitmq rabbitmqctl delete_queue events

docker-remove-rabbitmq:
	@docker container rm -f rabbitmq

docker-rabbitmq2:
	@docker run -d --name rabbitmq -o 5672:5672 -p 15672:15672 -e RABBITMQ_SERVER_ADDITIONAL_ERL_ARGS='-rabbitmq_stream advertised_host localhost -rabbit loopback_users "none"' rabbitmq:3.9-manegement

docker-rabbitmq-plugin-enable:
	@docker exec rabbitmq rabbitmq-plugins enable rabbitmq_stream
	@docker exec rabbitmq rabbitmq-plugins enable rabbitmq_stream_management

docker-tcpdump:
	@sudo tcpdump - i docker0 port 5552 -w rabbitmq.pcap

docker-tcpdump:
	@sudo tcpdump - i docker0 port 5552 -w rabbitmq_subentry.pcap

tcpdump-more:
	@tcpdump -A -r rabbitmq.pcap | more

tcpdump-sum:
	@tcpdump -r rabbitmq_subentry.pcap | grep 'length' | awk '{sum} {print sum}'
