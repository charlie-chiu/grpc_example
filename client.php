<?php

require_once "vendor/autoload.php";

use Greeting\GreeterClient;
use Greeting\HelloReply;
use Greeting\HelloRequest;

const addr = "host.docker.internal:9999";

$client = new GreeterClient(addr,
    ['credentials' => Grpc\ChannelCredentials::createInsecure()],
);

$request = new HelloRequest();
$request->setName("Charlie the PHPer");


/** @var HelloReply $response */
[$response, $status] = $client->SayHello($request)->wait();

echo sprintf("Response from greeter \"%s\", msg: \"%s\"\n", $response->getGreeter(), $response->getMessage());
