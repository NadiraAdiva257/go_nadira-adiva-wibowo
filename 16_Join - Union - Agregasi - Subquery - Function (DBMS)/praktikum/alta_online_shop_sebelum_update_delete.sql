-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 18, 2023 at 05:15 AM
-- Server version: 10.4.27-MariaDB
-- PHP Version: 8.2.0

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `alta_online_shop`
--

-- --------------------------------------------------------

--
-- Table structure for table `address`
--

CREATE TABLE `address` (
  `id` int(10) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `id_users` int(10) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- --------------------------------------------------------

--
-- Table structure for table `payment_methods`
--

CREATE TABLE `payment_methods` (
  `id` int(10) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `status` int(10) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `payment_methods`
--

INSERT INTO `payment_methods` (`id`, `name`, `status`, `description`, `created_at`, `updated_at`) VALUES
(1, 'transfer bank mandiri', NULL, 'transfer online atau offline melalui bank mandiri', '2023-03-17 04:13:33', '2023-03-17 04:13:33'),
(2, 'dana', NULL, 'pembayaran digital', '2023-03-17 04:13:33', '2023-03-17 04:13:33'),
(3, 'ovo', NULL, 'pembayaran digital', '2023-03-17 04:13:33', '2023-03-17 04:13:33');

-- --------------------------------------------------------

--
-- Table structure for table `products`
--

CREATE TABLE `products` (
  `id` int(10) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `description` varchar(255) DEFAULT NULL,
  `price` int(10) DEFAULT NULL,
  `id_product_operators` int(10) DEFAULT NULL,
  `id_product_types` int(10) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `products`
--

INSERT INTO `products` (`id`, `name`, `description`, `price`, `id_product_operators`, `id_product_types`, `created_at`, `updated_at`) VALUES
(1, 'notebook spiral', 'diimpor dari korea dengan desain yang unik', 36000, 3, 1, '2023-03-17 17:56:15', '2023-03-17 17:56:15'),
(2, 'rak mini', 'memiliki 9 laci transparan dengan warna yang soft', 25000, 3, 1, '2023-03-17 17:56:15', '2023-03-17 17:56:15'),
(5, 'seola senakers', 'bahan kulit sintetis dengan tinggi sol 4 cm', 170000, 1, 2, '2023-03-17 03:59:24', '2023-03-17 03:59:24'),
(6, 'kuki sneakers', 'bahan kanvas dengan tinggi sol 3 cm', 180000, 1, 2, '2023-03-17 03:59:24', '2023-03-17 03:59:24'),
(7, 'joonie sneakers', 'bahan kulit sintetis dengan tinggi sol 5 cm', 200000, 1, 2, '2023-03-17 03:59:24', '2023-03-17 03:59:24'),
(11, 'milo cereal box 150 gr', 'cereal dengan gandum utuh rasa coklat dan malt', 22000, 4, 3, '2023-03-17 04:07:58', '2023-03-17 04:07:58'),
(12, 'nescafe kaleng ice black 220 ml', 'minuman kopi siap minum dengan rasa yang mantap', 7000, 4, 3, '2023-03-17 04:07:58', '2023-03-17 04:07:58'),
(13, 'bear brand kaleng 189 ml', 'minuman susu siap minum dengan nutrisi susu sapi murni', 10000, 4, 3, '2023-03-17 04:07:58', '2023-03-17 04:07:58');

-- --------------------------------------------------------

--
-- Table structure for table `product_operators`
--

CREATE TABLE `product_operators` (
  `id` int(10) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `product_operators`
--

INSERT INTO `product_operators` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'pvn shoes', '2023-03-16 10:07:57', '2023-03-16 10:07:57'),
(2, 'kronikel', '2023-03-16 10:07:57', '2023-03-16 10:07:57'),
(3, 'miloli', '2023-03-16 10:07:57', '2023-03-16 10:07:57'),
(4, 'nestle indonesia', '2023-03-16 10:07:57', '2023-03-16 10:07:57'),
(5, 'kedai mart', '2023-03-16 10:07:57', '2023-03-16 10:07:57');

-- --------------------------------------------------------

--
-- Table structure for table `product_types`
--

CREATE TABLE `product_types` (
  `id` int(10) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `product_types`
--

INSERT INTO `product_types` (`id`, `name`, `created_at`, `updated_at`) VALUES
(1, 'alat tulis kantor', '2023-03-17 03:30:47', '2023-03-17 03:30:47'),
(2, 'sepatu wanita', '2023-03-17 03:30:47', '2023-03-17 03:30:47'),
(3, 'makanan minuman', '2023-03-17 03:30:47', '2023-03-17 03:30:47');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int(10) NOT NULL,
  `total_number_of_products` int(10) DEFAULT NULL,
  `total_price` int(10) DEFAULT NULL,
  `id_users` int(10) DEFAULT NULL,
  `id_payment_methods` int(10) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `total_number_of_products`, `total_price`, `id_users`, `id_payment_methods`, `created_at`, `updated_at`) VALUES
(1, 5, 86000, 10, 1, '2023-03-17 07:15:12', '2023-03-17 07:15:12'),
(2, 3, 216000, 10, 2, '2023-03-17 07:15:12', '2023-03-17 07:15:12'),
(3, 7, 231000, 10, 1, '2023-03-17 07:15:12', '2023-03-17 07:15:12'),
(4, 9, 111000, 11, 3, '2023-03-17 07:16:36', '2023-03-17 07:16:36'),
(5, 5, 298000, 11, 2, '2023-03-17 07:16:36', '2023-03-17 07:16:36'),
(6, 7, 82000, 11, 2, '2023-03-17 07:16:36', '2023-03-17 07:16:36'),
(7, 3, 372000, 13, 3, '2023-03-17 07:18:00', '2023-03-17 07:18:00'),
(8, 9, 191000, 13, 1, '2023-03-17 07:18:00', '2023-03-17 07:18:00'),
(9, 5, 104000, 13, 1, '2023-03-17 07:18:00', '2023-03-17 07:18:00'),
(10, 5, 257000, 14, 1, '2023-03-17 07:19:04', '2023-03-17 07:19:04'),
(11, 4, 76000, 14, 2, '2023-03-17 07:19:04', '2023-03-17 07:19:04'),
(12, 5, 298000, 14, 1, '2023-03-17 07:19:04', '2023-03-17 07:19:04'),
(13, 3, 39000, 12, 2, '2023-03-17 07:20:55', '2023-03-17 07:20:55'),
(14, 3, 228000, 12, 3, '2023-03-17 07:20:55', '2023-03-17 07:20:55'),
(15, 8, 86000, 12, 3, '2023-03-17 07:20:55', '2023-03-17 07:20:55');

-- --------------------------------------------------------

--
-- Table structure for table `transaction_details`
--

CREATE TABLE `transaction_details` (
  `id` int(10) NOT NULL,
  `id_transactions` int(10) NOT NULL,
  `id_products` int(10) NOT NULL,
  `number_of_products` int(10) DEFAULT NULL,
  `price` int(10) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `transaction_details`
--

INSERT INTO `transaction_details` (`id`, `id_transactions`, `id_products`, `number_of_products`, `price`, `created_at`, `updated_at`) VALUES
(1, 1, 2, 2, 25000, '2023-03-17 07:37:57', '2023-03-17 07:37:57'),
(2, 1, 11, 1, 22000, '2023-03-17 07:37:57', '2023-03-17 07:37:57'),
(3, 1, 12, 2, 7000, '2023-03-18 02:29:14', '2023-03-18 02:29:14'),
(4, 2, 5, 1, 170000, '2023-03-17 07:39:16', '2023-03-17 07:39:16'),
(5, 2, 13, 1, 10000, '2023-03-17 07:39:16', '2023-03-17 07:39:16'),
(6, 2, 1, 1, 36000, '2023-03-17 07:39:16', '2023-03-17 07:39:16'),
(7, 3, 6, 1, 180000, '2023-03-17 07:40:15', '2023-03-17 07:40:15'),
(8, 3, 12, 3, 7000, '2023-03-17 07:40:15', '2023-03-17 07:40:15'),
(9, 3, 13, 3, 10000, '2023-03-17 07:40:15', '2023-03-17 07:40:15'),
(10, 4, 11, 3, 22000, '2023-03-17 07:49:13', '2023-03-17 07:49:13'),
(11, 4, 12, 5, 7000, '2023-03-17 07:49:13', '2023-03-17 07:49:13'),
(12, 4, 13, 1, 10000, '2023-03-17 07:49:13', '2023-03-17 07:49:13'),
(13, 5, 6, 1, 180000, '2023-03-17 07:49:52', '2023-03-17 07:49:52'),
(14, 5, 1, 3, 36000, '2023-03-17 07:49:52', '2023-03-17 07:49:52'),
(15, 5, 13, 1, 10000, '2023-03-17 07:49:52', '2023-03-17 07:49:52'),
(16, 6, 12, 5, 7000, '2023-03-17 07:50:44', '2023-03-17 07:50:44'),
(17, 6, 2, 1, 25000, '2023-03-17 07:50:44', '2023-03-17 07:50:44'),
(18, 6, 11, 1, 22000, '2023-03-17 07:50:44', '2023-03-17 07:50:44'),
(19, 7, 5, 1, 170000, '2023-03-17 07:52:36', '2023-03-17 07:52:36'),
(20, 7, 6, 1, 180000, '2023-03-17 07:52:36', '2023-03-17 07:52:36'),
(21, 7, 11, 1, 22000, '2023-03-17 07:52:36', '2023-03-17 07:52:36'),
(22, 8, 13, 4, 10000, '2023-03-17 07:53:15', '2023-03-17 07:53:15'),
(23, 8, 12, 1, 7000, '2023-03-17 07:53:15', '2023-03-17 07:53:15'),
(24, 8, 1, 4, 36000, '2023-03-17 07:53:15', '2023-03-17 07:53:15'),
(25, 9, 2, 2, 25000, '2023-03-17 07:54:13', '2023-03-17 07:54:13'),
(26, 9, 11, 2, 22000, '2023-03-17 07:54:13', '2023-03-17 07:54:13'),
(27, 9, 13, 1, 10000, '2023-03-17 07:54:13', '2023-03-17 07:54:13'),
(28, 10, 12, 3, 7000, '2023-03-17 08:08:28', '2023-03-17 08:08:28'),
(29, 10, 7, 1, 200000, '2023-03-17 08:08:28', '2023-03-17 08:08:28'),
(30, 10, 1, 1, 36000, '2023-03-17 08:08:28', '2023-03-17 08:08:28'),
(31, 11, 2, 1, 25000, '2023-03-17 08:09:13', '2023-03-17 08:09:13'),
(32, 11, 11, 2, 22000, '2023-03-17 08:09:13', '2023-03-17 08:09:13'),
(33, 11, 12, 1, 7000, '2023-03-17 08:09:13', '2023-03-17 08:09:13'),
(34, 12, 6, 1, 180000, '2023-03-17 08:09:49', '2023-03-17 08:09:49'),
(35, 12, 1, 3, 36000, '2023-03-17 08:09:49', '2023-03-17 08:09:49'),
(36, 12, 13, 1, 10000, '2023-03-17 08:09:49', '2023-03-17 08:09:49'),
(37, 13, 11, 1, 22000, '2023-03-17 08:12:07', '2023-03-17 08:12:07'),
(38, 13, 12, 1, 7000, '2023-03-17 08:12:07', '2023-03-17 08:12:07'),
(39, 13, 13, 1, 10000, '2023-03-17 08:12:07', '2023-03-17 08:12:07'),
(40, 14, 1, 1, 36000, '2023-03-17 08:12:47', '2023-03-17 08:12:47'),
(41, 14, 5, 1, 170000, '2023-03-17 08:12:47', '2023-03-17 08:12:47'),
(42, 14, 11, 1, 22000, '2023-03-17 08:12:47', '2023-03-17 08:12:47'),
(43, 15, 12, 3, 7000, '2023-03-17 08:13:39', '2023-03-17 08:13:39'),
(44, 15, 13, 4, 10000, '2023-03-17 08:13:39', '2023-03-17 08:13:39'),
(45, 15, 2, 1, 25000, '2023-03-17 08:13:39', '2023-03-17 08:13:39');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(10) NOT NULL,
  `name` varchar(255) DEFAULT NULL,
  `birth_of_date` date DEFAULT NULL,
  `gender` char(1) DEFAULT NULL,
  `status` int(10) DEFAULT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `birth_of_date`, `gender`, `status`, `created_at`, `updated_at`) VALUES
(10, 'nadira', '2001-07-25', 'p', NULL, '2023-03-17 04:15:55', '2023-03-17 04:15:55'),
(11, 'lanang', '2002-12-18', 'l', NULL, '2023-03-16 04:19:49', '2023-03-17 04:15:55'),
(12, 'dimas', '2001-12-31', 'l', NULL, '2023-03-17 04:15:55', '2023-03-17 04:15:55'),
(13, 'sonya', '2001-03-14', 'p', NULL, '2023-03-17 04:15:55', '2023-03-17 04:15:55'),
(14, 'mayla', '2002-05-07', 'p', NULL, '2023-03-17 04:15:55', '2023-03-17 04:15:55');

-- --------------------------------------------------------

--
-- Table structure for table `user_payment_method_detail`
--

CREATE TABLE `user_payment_method_detail` (
  `id` int(10) NOT NULL,
  `id_users` int(10) NOT NULL,
  `id_payment_methods` int(10) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NOT NULL DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `address`
--
ALTER TABLE `address`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_users` (`id_users`);

--
-- Indexes for table `payment_methods`
--
ALTER TABLE `payment_methods`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `products`
--
ALTER TABLE `products`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_product_operators` (`id_product_operators`),
  ADD KEY `id_product_types` (`id_product_types`);

--
-- Indexes for table `product_operators`
--
ALTER TABLE `product_operators`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `product_types`
--
ALTER TABLE `product_types`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_users` (`id_users`),
  ADD KEY `id_payment_methods` (`id_payment_methods`);

--
-- Indexes for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_transactions` (`id_transactions`),
  ADD KEY `id_products` (`id_products`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `user_payment_method_detail`
--
ALTER TABLE `user_payment_method_detail`
  ADD PRIMARY KEY (`id`),
  ADD KEY `id_users` (`id_users`),
  ADD KEY `id_payment_methods` (`id_payment_methods`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `address`
--
ALTER TABLE `address`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `payment_methods`
--
ALTER TABLE `payment_methods`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `products`
--
ALTER TABLE `products`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=18;

--
-- AUTO_INCREMENT for table `product_operators`
--
ALTER TABLE `product_operators`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `product_types`
--
ALTER TABLE `product_types`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `transaction_details`
--
ALTER TABLE `transaction_details`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=46;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=15;

--
-- AUTO_INCREMENT for table `user_payment_method_detail`
--
ALTER TABLE `user_payment_method_detail`
  MODIFY `id` int(10) NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `address`
--
ALTER TABLE `address`
  ADD CONSTRAINT `address_ibfk_1` FOREIGN KEY (`id_users`) REFERENCES `users` (`id`);

--
-- Constraints for table `products`
--
ALTER TABLE `products`
  ADD CONSTRAINT `products_ibfk_1` FOREIGN KEY (`id_product_operators`) REFERENCES `product_operators` (`id`),
  ADD CONSTRAINT `products_ibfk_2` FOREIGN KEY (`id_product_types`) REFERENCES `product_types` (`id`);

--
-- Constraints for table `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `transactions_ibfk_1` FOREIGN KEY (`id_users`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `transactions_ibfk_2` FOREIGN KEY (`id_payment_methods`) REFERENCES `payment_methods` (`id`);

--
-- Constraints for table `transaction_details`
--
ALTER TABLE `transaction_details`
  ADD CONSTRAINT `transaction_details_ibfk_1` FOREIGN KEY (`id_transactions`) REFERENCES `transactions` (`id`),
  ADD CONSTRAINT `transaction_details_ibfk_2` FOREIGN KEY (`id_products`) REFERENCES `products` (`id`);

--
-- Constraints for table `user_payment_method_detail`
--
ALTER TABLE `user_payment_method_detail`
  ADD CONSTRAINT `user_payment_method_detail_ibfk_1` FOREIGN KEY (`id_users`) REFERENCES `users` (`id`),
  ADD CONSTRAINT `user_payment_method_detail_ibfk_2` FOREIGN KEY (`id_payment_methods`) REFERENCES `payment_methods` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
