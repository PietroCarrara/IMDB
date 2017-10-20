<?php
$servername = "localhost";
$username = "";
$password = "";
$dbname = "";

$db = new mysqli($servername, $username, $password, $dbname);

if ($db->connect_error) {
    die("Connection to database failed. Please contact the server owner.");
}
?>
