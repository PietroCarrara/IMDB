<?php

class Usuario {
	public usuario; // String
	private password; // String
	public id; // Int
	public watchlist; // Array de filmes
	public avaliados; // Associative array de [FILME] => float
	
	public passCheck($pass) {
		return $pass == $password;
	}
?>
