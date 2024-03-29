CREATE TABLE IF NOT EXISTS `tb_student` (
  `id` varchar(5) NOT NULL,
  `name` varchar(255) NOT NULL,
  `age` int(11) NOT NULL,
  `grade` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--@block
INSERT INTO `tb_student` (`id`, `name`, `age`, `grade`) VALUES
('B001', 'Jason Bourne', 29, 1),
('B002', 'James Bond', 27, 1),
('E001', 'Ethan Hunt', 27, 2),
('W001', 'John Wick', 28, 2);

--@block
ALTER TABLE `tb_student` ADD PRIMARY KEY (`id`);

--@block
SELECT * FROM tb_student;

--@block
DESCRIBE tb_student;

--@block
SHOW CREATE TABLE tb_student;

--@block
ALTER TABLE tb_student
ADD PRIMARY KEY (id);

--@block
select id, name, grade, age from tb_student where age > 27