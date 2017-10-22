<?php
$servername = "localhost";
$username = "IMDB_USER";
$password = "3T3Dp1uaNXAxbxWv";
$dbname = "IMDB";

$db = new mysqli($servername, $username, $password, $dbname);

if ($db->connect_error) {
    die("Connection to database failed. Please contact the server owner.");
}
?>
