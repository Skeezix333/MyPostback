<?php

$redis = new Redis(); 

try {
    $redis->connect('redis', getenv('REDIS_PORT'));
    $redis->auth(getenv('REDIS_PASSWORD'));
} catch (RedisException $e) {
    print_r($e);
}
echo "Connection to server sucessful. "; 
echo('<br>');
echo "Server is running: ".$redis->ping(); 
echo('<br>');

$json = file_get_contents("php://input");
$decoded = json_decode($json, true);

if ($decoded['data'] && $decoded['endpoint']) {
    $postback = $decoded["endpoint"];
    $postback["data"] = $decoded["data"][0];
    if ($postback) { 
        $redis->lPush('Post_Object', json_encode($postback));;
        echo('<br>');
        echo "Redis Push was successful!";
    }
}


$date = new DateTime();
$date = $date->format("y:m:d h:i:s");

?>