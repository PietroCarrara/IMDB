-- phpMyAdmin SQL Dump
-- version 4.7.5
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Nov 03, 2017 at 05:32 PM
-- Server version: 10.1.28-MariaDB
-- PHP Version: 7.1.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `IMDB`
--

-- --------------------------------------------------------

--
-- Table structure for table `cargo`
--

CREATE TABLE `cargo` (
  `ID` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `NOME` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `cargo`
--

INSERT INTO `cargo` (`ID`, `NOME`) VALUES
(NULL, 'Ator'),
(NULL, 'Diretor'),
(NULL, 'Sonoplasta'),
(NULL, 'Dublador');

-- --------------------------------------------------------

--
-- Table structure for table `filme`
--

CREATE TABLE `filme` (
  `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `sinopse` varchar(2048) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `filme`
--

INSERT INTO `filme` (`id`, `sinopse`) VALUES
(NULL, 'Ao viajar para a Califórnia para competir contra The King e Chick Hicks para o Campeonato da Copa Pistão, Lightning McQueen se perde depois de cair do seu caminhão na cidade chamada Radiator Springs. Aos poucos, ele faz amizade com os moradores da cidade, incluindo Sally, Doc Hudson e Mate. Quando chega a hora de McQueen ir embora, o campeonato não é mais a sua prioridade.'),
(NULL, 'Chihiro e seus pais estão se mudando para uma cidade diferente. A caminho da nova casa, o pai decide pegar um atalho. Eles se deparam com uma mesa repleta de comida, embora ninguém esteja por perto. Chihiro sente o perigo, mas seus pais começam a comer. Quando anoitece, eles se transformam em porcos. Agora, apenas Chihiro pode salvá-los.'),
(NULL, 'Don Cobb é um ladrão que invade os sonhos das pessoas e rouba segredos do subconsciente. As habilidades especiais de Cobb fazem com que ele seja procurado pelo mundo da espionagem empresarial, mas lhe custa tudo que ama. Cobb recebe uma missão impossível: plantar uma ideia na mente de uma pessoa. Se for bem-sucedido, será o crime perfeito, mas um amigo prevê todos os passos de Cobb.'),
(NULL, 'O mundo, em 2029, se tornou um local altamente informatizado, a ponto dos seres humanos poderem acessar extensas redes de informações com seu ciber-cérebros. A agente cibernética Major Motoko é a líder da unidade de serviço secreto Esquadrão Shell, responsável por combater o crime. Motoko foi tão modificada que quase todo seu corpo já é robótico. De humano só teria sobrado um fantasma de si mesma.');

-- --------------------------------------------------------

--
-- Table structure for table `filme_tag`
--

CREATE TABLE `filme_tag` (
  `ID_FILME` int(11) NOT NULL,
  `ID_TAG` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `filme_tag`
--

INSERT INTO `filme_tag` (`ID_FILME`, `ID_TAG`) VALUES
(1, 0),
(1, 2),
(1, 3),
(2, 3),
(3, 1),
(3, 2),
(4, 1),
(4, 2),
(4, 3);

-- --------------------------------------------------------

--
-- Table structure for table `imagem`
--

CREATE TABLE `imagem` (
  `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `id_filme` int(11) DEFAULT NULL,
  `caminho` varchar(256) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `imagem`
--

INSERT INTO `imagem` (`id`, `id_filme`, `caminho`) VALUES
(NULL, 3, '/uploads/1.jpg'),
(NULL, 3, '/res/img/screenshots/2/01.png'),
(NULL, 2, '/uploads/0.jpg'),
(NULL, 4, '/uploads/2.jpg'),
(NULL, 3, '/res/img/capas/3/capa.png'),
(NULL, 1, '/uploads/3.jpg');

-- --------------------------------------------------------

--
-- Table structure for table `lugar`
--

CREATE TABLE `lugar` (
  `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `nome` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `lugar`
--

INSERT INTO `lugar` (`id`, `nome`) VALUES
(NULL, 'Brasil'),
(NULL, 'Índia'),
(NULL, 'Estados Unidos'),
(NULL, 'México');

-- --------------------------------------------------------

--
-- Table structure for table `participa_como`
--

CREATE TABLE `participa_como` (
  `ID_FILME` int(11) NOT NULL,
  `ID_PESSOA` int(11) NOT NULL,
  `ID_CARGO` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `participa_como`
--

INSERT INTO `participa_como` (`ID_FILME`, `ID_PESSOA`, `ID_CARGO`) VALUES
(2, 5, 2),
(3, 1, 2),
(3, 2, 1),
(3, 3, 1),
(4, 4, 2);

-- --------------------------------------------------------

--
-- Table structure for table `pessoa`
--

CREATE TABLE `pessoa` (
  `ID` int(11) NOT NULL,
  `NOME` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `DATA_NASC` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `pessoa`
--

INSERT INTO `pessoa` (`ID`, `NOME`, `DATA_NASC`) VALUES
(NULL, 'Christopher Nolan', '1987-07-30'),
(NULL, 'Leonardo DiCaprio', '1974-11-11'),
(NULL, 'Ellen Page', '1987-02-21'),
(NULL, 'Shirow Masamune', '1963-11-23'),
(NULL, 'Rumi Hiiragi', '1987-08-01'),
(NULL, 'Teresa Eckton', '1950-05-06');

-- --------------------------------------------------------

--
-- Table structure for table `pessoa_filme`
--

CREATE TABLE `pessoa_filme` (
  `ID_FILME` int(11) NOT NULL,
  `ID_PESSOA` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `pessoa_filme`
--

INSERT INTO `pessoa_filme` (`ID_FILME`, `ID_PESSOA`) VALUES
(3, 1),
(3, 2),
(3, 3),
(4, 4),
(2, 5),
(1, 6);

-- --------------------------------------------------------

--
-- Table structure for table `pessoa_imagem`
--

CREATE TABLE `pessoa_imagem` (
  `id_pessoa` int(11) NOT NULL,
  `id_imagem` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `pessoa_imagem`
--

INSERT INTO `pessoa_imagem` (`id_pessoa`, `id_imagem`) VALUES
(2, 2),
(3, 2);

-- --------------------------------------------------------

--
-- Table structure for table `tag`
--

CREATE TABLE `tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `nome` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `tag`
--

INSERT INTO `tag` (`id`, `nome`) VALUES
(NULL, 'Terror'),
(NULL, 'Ficção Científica'),
(NULL, 'Ação'),
(NULL, 'Animação');

-- --------------------------------------------------------

--
-- Table structure for table `traducao`
--

CREATE TABLE `traducao` (
  `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `id_filme` int(11) NOT NULL,
  `ID_LUGAR` int(11) DEFAULT NULL,
  `DATA_LANC` date DEFAULT NULL,
  `TITULO` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `traducao`
--

INSERT INTO `traducao` (`id`, `id_filme`, `ID_LUGAR`, `DATA_LANC`, `TITULO`) VALUES
(NULL, 1, 1, '2000-04-13', 'Carros'),
(NULL, 2, 1, '2002-10-30', 'A Viagem De Chihiro'),
(NULL, 3, 1, '2010-08-10', 'A Origem'),
(NULL, 4, 1, '1996-03-29', 'O Fantasma do Futuro'),
(NULL, 1, 2, '2000-04-13', 'Cars'),
(NULL, 2, 3, '2002-04-20', 'Spirited Away'),
(NULL, 3, 2, '2010-07-16', 'Chakravuyh'),
(NULL, 4, 3, '1996-03-29', 'Ghost In The Shell'),
(NULL, 1, 3, '2000-04-13', 'Cars'),
(NULL, 2, 4, '2003-10-12', 'El viaje de Chihiro'),
(NULL, 3, 3, '2010-07-13', 'Inception'),
(NULL, 4, 2, '1996-03-29', 'Ghost in The Shell'),
(NULL, 1, 4, '2000-04-13', 'Cars'),
(NULL, 2, 2, '2002-04-20', 'Spirited Away'),
(NULL, 3, 4, '2010-07-10', 'El Origen');

-- --------------------------------------------------------

--
-- Table structure for table `usuario`
--

CREATE TABLE `usuario` (
  `id` int(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `nome` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `senha` varchar(1024) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `is_admin` tinyint(4) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `usuario`
--

INSERT INTO `usuario` (`id`, `nome`, `senha`, `is_admin`) VALUES
(NULL, 'pietro', 'senha', 0),
(NULL, 'arthur', 'senha', 0),
(NULL, 'root', 'root', 1);

-- --------------------------------------------------------

--
-- Table structure for table `usuario_filme_lista`
--

CREATE TABLE `usuario_filme_lista` (
  `ID_FILME` int(11) NOT NULL,
  `ID_USUARIO` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `usuario_filme_lista`
--

INSERT INTO `usuario_filme_lista` (`ID_FILME`, `ID_USUARIO`) VALUES
(1, 1),
(3, 2),
(4, 1),
(4, 2);

-- --------------------------------------------------------

--
-- Table structure for table `usuario_filme_nota`
--

CREATE TABLE `usuario_filme_nota` (
  `ID_FILME` int(11) NOT NULL,
  `ID_USUARIO` int(11) NOT NULL,
  `nota` int(11) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `usuario_filme_nota`
--

INSERT INTO `usuario_filme_nota` (`ID_FILME`, `ID_USUARIO`, `nota`) VALUES
(1, 1, 5),
(2, 1, 4),
(2, 2, 5);

-- --------------------------------------------------------

--
-- Table structure for table `usuario_filme_review`
--

CREATE TABLE `usuario_filme_review` (
  `ID` int(11) AUTO_INCREMENT,
  `ID_FILME` int(11) NOT NULL,
  `ID_USUARIO` int(11) NOT NULL,
  `TEXTO` varchar(512) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

--
-- Dumping data for table `usuario_filme_review`
--

INSERT INTO `usuario_filme_review` (`ID_FILME`, `ID_USUARIO`, `TEXTO`) VALUES
(1, 1, 'Bah bem tri esse filme daí'),
(1, 2, 'Curti pra caramba'),
(4, 1, 'Esse eu recomendo'),
(4, 2, 'Esse foi bem massa');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `filme_tag`
--
ALTER TABLE `filme_tag`
  ADD PRIMARY KEY (`ID_FILME`,`ID_TAG`),
  ADD KEY `FK_FILME_TAG_TAG` (`ID_TAG`);

--
-- Indexes for table `imagem`
--
ALTER TABLE `imagem`
  ADD KEY `id_filme` (`id_filme`);

--
-- Indexes for table `participa_como`
--
ALTER TABLE `participa_como`
  ADD PRIMARY KEY (`ID_FILME`,`ID_PESSOA`,`ID_CARGO`),
  ADD KEY `FK_PARTICIPA_COMO_PESSOA` (`ID_PESSOA`),
  ADD KEY `FK_PARTICIPA_COMO_CARGO` (`ID_CARGO`);

--
-- Indexes for table `pessoa_filme`
--
ALTER TABLE `pessoa_filme`
  ADD PRIMARY KEY (`ID_PESSOA`,`ID_FILME`),
  ADD KEY `FK_PESSOA_FILME_FILME` (`ID_FILME`);

--
-- Indexes for table `pessoa_imagem`
--
ALTER TABLE `pessoa_imagem`
  ADD PRIMARY KEY (`id_pessoa`,`id_imagem`),
  ADD KEY `FK_PESSOA_IMAGEM_IMAGEM` (`id_imagem`);

--
-- Indexes for table `usuario_filme_lista`
--
ALTER TABLE `usuario_filme_lista`
  ADD PRIMARY KEY (`ID_FILME`,`ID_USUARIO`),
  ADD KEY `FK_USUARIO_FILME_LISTA_USUARIO` (`ID_USUARIO`);

--
-- Indexes for table `usuario_filme_nota`
--
ALTER TABLE `usuario_filme_nota`
  ADD PRIMARY KEY (`ID_FILME`,`ID_USUARIO`),
  ADD KEY `FK_USUARIO_FILME_NOTA_USUARIO` (`ID_USUARIO`);

--
-- Indexes for table `usuario_filme_review`
--
ALTER TABLE `usuario_filme_review`
  --ADD PRIMARY KEY (`ID_FILME`,`ID_USUARIO`),
  ADD PRIMARY KEY (`ID`),
  ADD KEY `FK_USUARIO_FILME_REVIEW_USUARIO` (`ID_USUARIO`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `filme_tag`
--
-- Essas falham, não sei pq
-- Apesar disso, elas não necessárias
-- para o funcionamento do banco :^)
ALTER TABLE `filme_tag`
  ADD CONSTRAINT `FK_FILME_TAG_FILME` FOREIGN KEY (`ID_FILME`) REFERENCES `filme` (`id`),
  ADD CONSTRAINT `FK_FILME_TAG_TAG` FOREIGN KEY (`ID_TAG`) REFERENCES `tag` (`id`);

--
-- Constraints for table `imagem`
--
ALTER TABLE `imagem`
  ADD CONSTRAINT `FK_IMAGEM_FILME` FOREIGN KEY (`id_filme`) REFERENCES `filme` (`id`);

--
-- Constraints for table `participa_como`
--
ALTER TABLE `participa_como`
  ADD CONSTRAINT `FK_PARTICIPA_COMO_CARGO` FOREIGN KEY (`ID_CARGO`) REFERENCES `cargo` (`ID`),
  ADD CONSTRAINT `FK_PARTICIPA_COMO_FILME` FOREIGN KEY (`ID_FILME`) REFERENCES `filme` (`id`),
  ADD CONSTRAINT `FK_PARTICIPA_COMO_PESSOA` FOREIGN KEY (`ID_PESSOA`) REFERENCES `pessoa` (`ID`);

--
-- Constraints for table `pessoa_filme`
--
ALTER TABLE `pessoa_filme`
  ADD CONSTRAINT `FK_PESSOA_FILME_FILME` FOREIGN KEY (`ID_FILME`) REFERENCES `filme` (`id`),
  ADD CONSTRAINT `FK_PESSOA_FILME_PESSOA` FOREIGN KEY (`ID_PESSOA`) REFERENCES `pessoa` (`ID`);

--
-- Constraints for table `pessoa_imagem`
--
ALTER TABLE `pessoa_imagem`
  ADD CONSTRAINT `FK_PESSOA_IMAGEM_IMAGEM` FOREIGN KEY (`id_imagem`) REFERENCES `imagem` (`id`),
  ADD CONSTRAINT `FK_PESSOA_IMAGEM_PESSOA` FOREIGN KEY (`id_pessoa`) REFERENCES `pessoa` (`ID`);

--
-- Constraints for table `traducao`
--
ALTER TABLE `traducao`
  ADD CONSTRAINT `FK_TRADUCAO_FILME` FOREIGN KEY (`id_filme`) REFERENCES `filme` (`id`),
  ADD CONSTRAINT `FK_TRADUCAO_LUGAR` FOREIGN KEY (`ID_LUGAR`) REFERENCES `lugar` (`id`);

--
-- Constraints for table `usuario_filme_lista`
--
ALTER TABLE `usuario_filme_lista`
  ADD CONSTRAINT `FK_USUARIO_FILME_LISTA_FILME` FOREIGN KEY (`ID_FILME`) REFERENCES `filme` (`id`),
  ADD CONSTRAINT `FK_USUARIO_FILME_LISTA_USUARIO` FOREIGN KEY (`ID_USUARIO`) REFERENCES `usuario` (`id`);

--
-- Constraints for table `usuario_filme_nota`
--
ALTER TABLE `usuario_filme_nota`
  ADD CONSTRAINT `FK_USUARIO_FILME_NOTA_FILME` FOREIGN KEY (`ID_FILME`) REFERENCES `filme` (`id`),
  ADD CONSTRAINT `FK_USUARIO_FILME_NOTA_USUARIO` FOREIGN KEY (`ID_USUARIO`) REFERENCES `usuario` (`id`);

--
-- Constraints for table `usuario_filme_review`
--
ALTER TABLE `usuario_filme_review`
  ADD CONSTRAINT `FK_USUARIO_FILME_REVIEW_FILME` FOREIGN KEY (`ID_FILME`) REFERENCES `filme` (`id`),
  ADD CONSTRAINT `FK_USUARIO_FILME_REVIEW_USUARIO` FOREIGN KEY (`ID_USUARIO`) REFERENCES `usuario` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
