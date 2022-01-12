<?php
$start = microtime(true);

$result = 0;
for ($i = 0; $i < 10 ** 8; $i++) {
    $result += $i;
}

$time = microtime(true) - $start;

echo "time: ${time}[s]", PHP_EOL, "result: $result", PHP_EOL;
