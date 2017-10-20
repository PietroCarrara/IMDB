<?php
include 'db.php';
include '../../vendor/autoload.php';

$m = new Mustache_Engine(array(
    'loader' => new Mustache_Loader_FilesystemLoader('/app/view/'),
));

?>
