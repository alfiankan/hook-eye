<?php

$conf = new RdKafka\Conf();
$conf->set('metadata.broker.list', '18.139.110.18:9092');
$message = "hell yeah";
//If you need to produce exactly once and want to keep the original produce order, uncomment the line below
//$conf->set('enable.idempotence', 'true');

$producer = new RdKafka\Producer($conf);

$topic = $producer->newTopic("payment");
$topic->produce(RD_KAFKA_PARTITION_UA, 0, $message);

$producer->poll(0);
$result = $producer->flush(10000);
if (RD_KAFKA_RESP_ERR_NO_ERROR === $result) {
    return true;
} else {
    print_r("error");
}