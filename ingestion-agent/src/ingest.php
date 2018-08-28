<?php

$redis = new Redis(); 

try {
    $redis->connect('redis', getenv('REDIS_PORT'));
    $redis->auth(getenv('REDIS_PASSWORD'));
} catch (RedisException $e) {
    print_r($e);
}
echo "Connection to server sucessful. "; 
echo "Server is running: ".$redis->ping(); 

$json = file_get_contents("php://input");
$decoded = json_decode($json, true);

if ($decoded['data'] && $decoded['endpoint']) {
    $postback = array(
        "endpoint" => $decoded['endpoint'],
        "data" => $decoded['data'],
    );
    $redis->lPush('Post_Object', json_encode($postback));
    echo "Redis Push was successful!";
}


$date = new DateTime();
$date = $date->format("y:m:d h:i:s");


echo($decoded['data']);
echo($postback['data']);
echo($postback['endpoint']);

?>