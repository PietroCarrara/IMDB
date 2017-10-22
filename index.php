<?php

echo 'sla1';

require_once('config/conf.php');

echo 'sla';

$template = file_get_contents('app/view/index.html');

echo $m->render($template, array('titulo' => 'World!')); // "Hello World!"

?>
