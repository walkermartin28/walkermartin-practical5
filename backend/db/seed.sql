-- SEC Baseball Hub Database Seed Data
-- Populates database with realistic 2025 season data
-- Database: sec_baseball

USE sec_baseball;

-- ============================================================
-- Insert Teams (All 16 SEC Baseball Teams)
-- ============================================================

INSERT INTO teams (name, abbreviation, logo_url, home_stadium, location, conference_wins, conference_losses, overall_wins, overall_losses, runs_scored, runs_allowed) VALUES
('Alabama Crimson Tide', 'ALA', 'https://a.espncdn.com/i/teamlogos/ncaa/500/333.png', 'Sewell-Thomas Stadium', 'Tuscaloosa, AL', 12, 6, 28, 12, 245, 198),
('Arkansas Razorbacks', 'ARK', 'https://a.espncdn.com/i/teamlogos/ncaa/500/8.png', 'Baum-Walker Stadium', 'Fayetteville, AR', 15, 3, 32, 8, 298, 167),
('Auburn Tigers', 'AUB', 'https://a.espncdn.com/i/teamlogos/ncaa/500/2.png', 'Plainsman Park', 'Auburn, AL', 10, 8, 25, 15, 221, 205),
('Florida Gators', 'FLA', 'https://a.espncdn.com/i/teamlogos/ncaa/500/57.png', 'Florida Ballpark', 'Gainesville, FL', 14, 4, 31, 9, 276, 178),
('Georgia Bulldogs', 'UGA', 'https://a.espncdn.com/i/teamlogos/ncaa/500/61.png', 'Foley Field', 'Athens, GA', 11, 7, 27, 13, 238, 192),
('Kentucky Wildcats', 'UK', 'https://a.espncdn.com/i/teamlogos/ncaa/500/96.png', 'Kentucky Proud Park', 'Lexington, KY', 8, 10, 22, 18, 201, 223),
('LSU Tigers', 'LSU', 'https://a.espncdn.com/i/teamlogos/ncaa/500/99.png', 'Alex Box Stadium', 'Baton Rouge, LA', 13, 5, 30, 10, 287, 185),
('Mississippi State Bulldogs', 'MSST', 'https://a.espncdn.com/i/teamlogos/ncaa/500/344.png', 'Dudy Noble Field', 'Starkville, MS', 9, 9, 24, 16, 215, 210),
('Missouri Tigers', 'MIZ', 'https://a.espncdn.com/i/teamlogos/ncaa/500/142.png', 'Taylor Stadium', 'Columbia, MO', 7, 11, 20, 20, 189, 241),
('Ole Miss Rebels', 'MISS', 'https://a.espncdn.com/i/teamlogos/ncaa/500/145.png', 'Swayze Field', 'Oxford, MS', 11, 7, 26, 14, 252, 201),
('South Carolina Gamecocks', 'SC', 'https://a.espncdn.com/i/teamlogos/ncaa/500/2579.png', 'Founders Park', 'Columbia, SC', 9, 9, 23, 17, 208, 217),
('Tennessee Volunteers', 'TENN', 'https://a.espncdn.com/i/teamlogos/ncaa/500/2633.png', 'Lindsey Nelson Stadium', 'Knoxville, TN', 16, 2, 34, 6, 312, 156),
('Texas A&M Aggies', 'TAMU', 'https://a.espncdn.com/i/teamlogos/ncaa/500/245.png', 'Blue Bell Park', 'College Station, TX', 12, 6, 29, 11, 267, 189),
('Vanderbilt Commodores', 'VU', 'https://a.espncdn.com/i/teamlogos/ncaa/500/238.png', 'Hawkins Field', 'Nashville, TN', 10, 8, 26, 14, 233, 203),
('Texas Longhorns', 'TEX', 'https://a.espncdn.com/i/teamlogos/ncaa/500/251.png', 'UFCU Disch-Falk Field', 'Austin, TX', 13, 5, 31, 9, 289, 172),
('Oklahoma Sooners', 'OU', 'https://a.espncdn.com/i/teamlogos/ncaa/500/201.png', 'L. Dale Mitchell Park', 'Norman, OK', 11, 7, 27, 13, 241, 195);

-- ============================================================
-- Insert Players (6 players per team = 96 total)
-- Mix of pitchers and position players with realistic stats
-- ============================================================

-- Alabama Players
INSERT INTO players (team_id, first_name, last_name, jersey_number, position, year, height, weight, batting_avg, home_runs, rbis, era, strikeouts, wins, saves, innings_pitched, games_played) VALUES
(1, 'Tommy', 'Seidl', 8, 'SS', 'Junior', '6-0', 185, 0.312, 7, 38, NULL, 42, 0, 0, NULL, 40),
(1, 'Bryce', 'Eblin', 25, 'OF', 'Senior', '6-2', 200, 0.298, 11, 45, NULL, 38, 0, 0, NULL, 40),
(1, 'Luke', 'Holman', 33, 'P', 'Sophomore', '6-4', 215, NULL, 0, 0, 3.12, 78, 6, 0, 72.0, 14),
(1, 'Ben', 'Hess', 16, 'C', 'Junior', '6-1', 210, 0.285, 9, 41, NULL, 51, 0, 0, NULL, 39),
(1, 'Gage', 'Ragland', 11, 'P', 'Senior', '6-3', 205, NULL, 0, 0, 2.89, 65, 7, 0, 68.1, 13),
(1, 'Drew', 'Williamson', 3, '2B', 'Freshman', '5-11', 180, 0.267, 3, 24, NULL, 28, 0, 0, NULL, 35),

-- Arkansas Players
(2, 'Peyton', 'Holt', 1, 'SS', 'Junior', '6-0', 190, 0.338, 9, 52, NULL, 35, 0, 0, NULL, 40),
(2, 'Wehiwa', 'Aloy', 24, 'OF', 'Senior', '6-3', 215, 0.321, 14, 58, NULL, 47, 0, 0, NULL, 40),
(2, 'Hagen', 'Smith', 33, 'P', 'Junior', '6-3', 220, NULL, 0, 0, 2.45, 95, 8, 0, 84.2, 15),
(2, 'Parker', 'Rowland', 12, 'C', 'Sophomore', '6-2', 215, 0.305, 12, 49, NULL, 44, 0, 0, NULL, 38),
(2, 'Mason', 'Molina', 19, 'P', 'Senior', '6-4', 225, NULL, 0, 0, 2.78, 72, 7, 0, 71.0, 14),
(2, 'Ben', 'McLaughlin', 7, '3B', 'Junior', '6-1', 200, 0.294, 8, 43, NULL, 39, 0, 0, NULL, 39),

-- Auburn Players
(3, 'Cooper', 'McMurray', 5, 'C', 'Junior', '6-0', 205, 0.289, 10, 44, NULL, 48, 0, 0, NULL, 40),
(3, 'Carter', 'Wright', 9, 'SS', 'Sophomore', '6-1', 185, 0.276, 5, 32, NULL, 33, 0, 0, NULL, 38),
(3, 'Chase', 'Allsup', 27, 'P', 'Senior', '6-2', 200, NULL, 0, 0, 3.54, 68, 5, 0, 66.0, 13),
(3, 'Bobby', 'Pierce', 21, 'OF', 'Junior', '6-3', 210, 0.302, 9, 39, NULL, 41, 0, 0, NULL, 39),
(3, 'Carson', 'Skipper', 31, 'P', 'Freshman', '6-5', 220, NULL, 0, 0, 3.21, 59, 5, 0, 61.2, 12),
(3, 'Ike', 'Irish', 4, '2B', 'Senior', '5-10', 175, 0.268, 4, 28, NULL, 29, 0, 0, NULL, 37),

-- Florida Players
(4, 'Jac', 'Caglianone', 35, 'P/1B', 'Senior', '6-5', 245, 0.328, 18, 62, 2.56, 82, 8, 0, 77.1, 40),
(4, 'Cade', 'Kurland', 10, 'SS', 'Junior', '6-2', 195, 0.315, 8, 47, NULL, 37, 0, 0, NULL, 40),
(4, 'Luke', 'Heyman', 28, 'P', 'Sophomore', '6-3', 210, NULL, 0, 0, 2.89, 76, 7, 0, 71.2, 14),
(4, 'Michael', 'Robertson', 8, 'OF', 'Senior', '6-1', 200, 0.307, 11, 51, NULL, 43, 0, 0, NULL, 39),
(4, 'Brody', 'Donay', 25, 'C', 'Junior', '6-0', 215, 0.291, 9, 45, NULL, 46, 0, 0, NULL, 38),
(4, 'Brandon', 'Neely', 32, 'P', 'Junior', '6-4', 215, NULL, 0, 0, 3.02, 64, 6, 1, 65.1, 18),

-- Georgia Players
(5, 'Charlie', 'Condon', 3, '3B', 'Junior', '6-6', 215, 0.334, 15, 56, NULL, 49, 0, 0, NULL, 40),
(5, 'Corey', 'Collins', 1, 'SS', 'Senior', '5-11', 180, 0.298, 6, 38, NULL, 35, 0, 0, NULL, 40),
(5, 'Leighton', 'Finley', 31, 'P', 'Junior', '6-6', 230, NULL, 0, 0, 3.18, 71, 6, 0, 68.0, 13),
(5, 'Fernando', 'Gonzalez', 9, 'C', 'Sophomore', '6-1', 210, 0.282, 8, 42, NULL, 44, 0, 0, NULL, 39),
(5, 'Kolten', 'Smith', 24, 'OF', 'Senior', '6-2', 195, 0.289, 7, 35, NULL, 38, 0, 0, NULL, 38),
(5, 'Chandler', 'Marsh', 29, 'P', 'Freshman', '6-3', 205, NULL, 0, 0, 3.45, 55, 5, 0, 60.0, 12),

-- Kentucky Players
(6, 'Ryan', 'Waldschmidt', 22, 'OF', 'Senior', '6-3', 210, 0.301, 12, 48, NULL, 52, 0, 0, NULL, 40),
(6, 'James', 'McCoy', 2, 'SS', 'Junior', '6-0', 185, 0.278, 5, 31, NULL, 34, 0, 0, NULL, 38),
(6, 'Mason', 'Moore', 27, 'P', 'Senior', '6-4', 215, NULL, 0, 0, 3.87, 61, 4, 0, 63.0, 13),
(6, 'Nolan', 'McCarthy', 33, 'C', 'Sophomore', '6-2', 220, 0.265, 7, 36, NULL, 48, 0, 0, NULL, 37),
(6, 'Trey', 'Pooser', 11, 'P', 'Junior', '6-1', 195, NULL, 0, 0, 4.12, 58, 4, 2, 59.0, 16),
(6, 'Grant', 'Smith', 7, '1B', 'Senior', '6-4', 225, 0.289, 9, 41, NULL, 39, 0, 0, NULL, 39),

-- LSU Players
(7, 'Tommy', 'White', 4, '3B', 'Junior', '5-10', 210, 0.325, 16, 61, NULL, 56, 0, 0, NULL, 40),
(7, 'Josh', 'Pearson', 18, 'OF', 'Senior', '6-4', 225, 0.318, 13, 54, NULL, 48, 0, 0, NULL, 40),
(7, 'Griffin', 'Herring', 38, 'P', 'Sophomore', '6-2', 200, NULL, 0, 0, 2.67, 89, 8, 0, 77.2, 14),
(7, 'Alex', 'Milazzo', 13, 'C', 'Junior', '6-1', 215, 0.295, 8, 46, NULL, 43, 0, 0, NULL, 39),
(7, 'Luke', 'Holman', 36, 'P', 'Junior', '6-5', 220, NULL, 0, 0, 2.91, 74, 7, 0, 71.0, 13),
(7, 'Steven', 'Milam', 5, 'SS', 'Sophomore', '6-0', 190, 0.287, 6, 39, NULL, 37, 0, 0, NULL, 38),

-- Mississippi State Players
(8, 'Dakota', 'Jordan', 8, 'SS', 'Senior', '6-1', 190, 0.293, 7, 37, NULL, 36, 0, 0, NULL, 40),
(8, 'Kellum', 'Clark', 9, '1B', 'Junior', '6-3', 220, 0.305, 11, 49, NULL, 47, 0, 0, NULL, 39),
(8, 'Jurrangelo', 'Cijntje', 23, 'P', 'Sophomore', '6-7', 245, NULL, 0, 0, 3.34, 72, 6, 0, 67.1, 13),
(8, 'David', 'Mershon', 44, 'C', 'Senior', '6-2', 215, 0.278, 8, 41, NULL, 45, 0, 0, NULL, 38),
(8, 'Brooks', 'Auger', 34, 'P', 'Junior', '6-4', 210, NULL, 0, 0, 3.76, 63, 5, 0, 64.2, 12),
(8, 'Amani', 'Larry', 3, 'OF', 'Freshman', '6-1', 195, 0.271, 5, 29, NULL, 31, 0, 0, NULL, 36),

-- Missouri Players
(9, 'Trey', 'Dillard', 2, 'SS', 'Junior', '6-0', 180, 0.281, 4, 27, NULL, 32, 0, 0, NULL, 38),
(9, 'Matt', 'Snider', 26, 'OF', 'Senior', '6-2', 200, 0.274, 6, 33, NULL, 38, 0, 0, NULL, 37),
(9, 'Ty', 'Buckner', 17, 'P', 'Sophomore', '6-3', 205, NULL, 0, 0, 4.21, 54, 4, 0, 62.0, 13),
(9, 'Jaden', 'Anderson', 15, 'C', 'Junior', '6-1', 210, 0.265, 5, 28, NULL, 41, 0, 0, NULL, 36),
(9, 'Hunter', 'Lasley', 22, 'P', 'Senior', '6-2', 195, NULL, 0, 0, 4.56, 49, 3, 1, 57.1, 15),
(9, 'Brayden', 'Williams', 10, '2B', 'Freshman', '5-11', 175, 0.258, 3, 22, NULL, 27, 0, 0, NULL, 34),

-- Ole Miss Players
(10, 'Ethan', 'Groff', 6, 'C', 'Junior', '6-2', 220, 0.309, 13, 53, NULL, 51, 0, 0, NULL, 40),
(10, 'Judd', 'Utermark', 14, 'SS', 'Senior', '6-1', 190, 0.298, 7, 42, NULL, 38, 0, 0, NULL, 39),
(10, 'Slade', 'Tepera', 30, 'P', 'Junior', '6-3', 210, NULL, 0, 0, 3.28, 69, 6, 0, 68.2, 13),
(10, 'Ethan', 'Lege', 9, 'OF', 'Sophomore', '6-0', 195, 0.287, 8, 39, NULL, 42, 0, 0, NULL, 38),
(10, 'Austin', 'Bost', 27, 'P', 'Senior', '6-4', 215, NULL, 0, 0, 3.51, 65, 6, 0, 66.2, 12),
(10, 'Will', 'Furniss', 3, '2B', 'Junior', '5-10', 180, 0.279, 4, 31, NULL, 29, 0, 0, NULL, 37),

-- South Carolina Players
(11, 'Cole', 'Messina', 19, 'SS', 'Junior', '6-1', 195, 0.291, 6, 36, NULL, 34, 0, 0, NULL, 39),
(11, 'Ethan', 'Petry', 44, 'OF', 'Senior', '6-3', 205, 0.284, 9, 41, NULL, 44, 0, 0, NULL, 38),
(11, 'Dylan', 'Brewer', 25, 'P', 'Sophomore', '6-2', 200, NULL, 0, 0, 3.62, 67, 5, 0, 64.1, 13),
(11, 'Talmadge', 'LeCroy', 8, 'C', 'Junior', '6-0', 210, 0.273, 7, 38, NULL, 46, 0, 0, NULL, 37),
(11, 'James', 'Hicks', 31, 'P', 'Junior', '6-5', 225, NULL, 0, 0, 3.89, 61, 5, 0, 62.1, 12),
(11, 'Kennedy', 'Jones', 2, '3B', 'Freshman', '6-2', 200, 0.268, 5, 29, NULL, 31, 0, 0, NULL, 36),

-- Tennessee Players
(12, 'Billy', 'Amick', 7, 'SS', 'Junior', '6-0', 185, 0.342, 10, 58, NULL, 39, 0, 0, NULL, 40),
(12, 'Hunter', 'Ensley', 31, 'OF', 'Senior', '6-4', 220, 0.335, 17, 64, NULL, 52, 0, 0, NULL, 40),
(12, 'Drew', 'Beam', 14, 'P', 'Junior', '6-2', 200, NULL, 0, 0, 2.21, 98, 9, 0, 85.1, 15),
(12, 'Cal', 'Stark', 5, 'C', 'Sophomore', '6-1', 215, 0.318, 12, 55, NULL, 47, 0, 0, NULL, 39),
(12, 'AJ', 'Causey', 20, 'P', 'Senior', '6-3', 210, NULL, 0, 0, 2.54, 81, 8, 0, 78.0, 14),
(12, 'Kavares', 'Tears', 22, '2B', 'Junior', '5-11', 180, 0.311, 6, 41, NULL, 34, 0, 0, NULL, 38),

-- Texas A&M Players
(13, 'Kaeden', 'Kent', 4, 'SS', 'Junior', '6-2', 195, 0.316, 9, 49, NULL, 38, 0, 0, NULL, 40),
(13, 'Jace', 'LaViolette', 17, '1B', 'Senior', '6-4', 230, 0.322, 15, 59, NULL, 54, 0, 0, NULL, 40),
(13, 'Ryan', 'Prager', 28, 'P', 'Junior', '6-1', 195, NULL, 0, 0, 2.87, 84, 8, 0, 75.2, 14),
(13, 'Jackson', 'Appel', 26, 'C', 'Sophomore', '6-2', 215, 0.298, 10, 46, NULL, 48, 0, 0, NULL, 39),
(13, 'Justin', 'Lamkin', 35, 'P', 'Senior', '6-3', 215, NULL, 0, 0, 3.11, 72, 7, 0, 72.0, 13),
(13, 'Gavin', 'Grahovac', 13, 'OF', 'Junior', '6-1', 200, 0.289, 8, 42, NULL, 41, 0, 0, NULL, 38),

-- Vanderbilt Players
(14, 'RJ', 'Schreck', 6, 'C', 'Junior', '6-3', 225, 0.304, 11, 48, NULL, 49, 0, 0, NULL, 40),
(14, 'Davis', 'Diaz', 3, 'SS', 'Sophomore', '6-1', 185, 0.292, 6, 37, NULL, 35, 0, 0, NULL, 39),
(14, 'Chris', 'McElvain', 32, 'P', 'Senior', '6-5', 220, NULL, 0, 0, 3.24, 75, 6, 0, 69.1, 13),
(14, 'Jonathan', 'Vastine', 18, 'OF', 'Junior', '6-2', 200, 0.287, 8, 41, NULL, 43, 0, 0, NULL, 38),
(14, 'Greysen', 'Carter', 23, 'P', 'Junior', '6-4', 210, NULL, 0, 0, 3.47, 68, 6, 0, 67.0, 12),
(14, 'Parker', 'Noland', 10, '3B', 'Senior', '6-0', 195, 0.281, 7, 36, NULL, 37, 0, 0, NULL, 37),

-- Texas Players
(15, 'Max', 'Belyeu', 19, 'OF', 'Junior', '6-1', 200, 0.331, 14, 57, NULL, 51, 0, 0, NULL, 40),
(15, 'Jalin', 'Flores', 0, 'SS', 'Senior', '5-11', 185, 0.319, 8, 48, NULL, 37, 0, 0, NULL, 40),
(15, 'Trey', 'Woosley', 27, 'P', 'Junior', '6-3', 210, NULL, 0, 0, 2.73, 87, 8, 0, 79.0, 14),
(15, 'Rylan', 'Galvan', 3, 'C', 'Sophomore', '6-2', 220, 0.307, 11, 52, NULL, 50, 0, 0, NULL, 39),
(15, 'Jared', 'Southard', 22, 'P', 'Senior', '6-4', 215, NULL, 0, 0, 2.98, 76, 7, 0, 72.1, 13),
(15, 'Will', 'Gasparino', 13, '2B', 'Junior', '6-0', 190, 0.295, 7, 43, NULL, 38, 0, 0, NULL, 38),

-- Oklahoma Players
(16, 'Kendall', 'Pettis', 6, 'SS', 'Senior', '6-1', 190, 0.308, 9, 45, NULL, 40, 0, 0, NULL, 40),
(16, 'Jaxon', 'Willits', 18, 'OF', 'Junior', '6-3', 210, 0.301, 10, 47, NULL, 44, 0, 0, NULL, 39),
(16, 'Braden', 'Carmichael', 24, 'P', 'Sophomore', '6-5', 225, NULL, 0, 0, 3.15, 73, 7, 0, 71.1, 13),
(16, 'Jimmy', 'Crooks', 9, 'C', 'Junior', '6-2', 215, 0.286, 8, 41, NULL, 46, 0, 0, NULL, 38),
(16, 'Jake', 'Bennett', 30, 'P', 'Senior', '6-2', 200, NULL, 0, 0, 3.38, 66, 6, 0, 66.2, 12),
(16, 'Easton', 'Carmichael', 2, '3B', 'Freshman', '6-1', 195, 0.277, 6, 35, NULL, 36, 0, 0, NULL, 37);

-- ============================================================
-- Insert Games (Mix of completed and scheduled games)
-- ============================================================

INSERT INTO games (home_team_id, away_team_id, game_date, game_time, location, status, home_score, away_score, is_conference_game, attendance) VALUES
-- Completed Games
(12, 14, '2025-03-15', '14:00:00', 'Lindsey Nelson Stadium', 'completed', 8, 3, TRUE, 4162),
(2, 7, '2025-03-16', '15:00:00', 'Baum-Walker Stadium', 'completed', 5, 6, TRUE, 11084),
(15, 13, '2025-03-22', '18:30:00', 'UFCU Disch-Falk Field', 'completed', 7, 4, TRUE, 7823),
(4, 5, '2025-03-23', '13:00:00', 'Florida Ballpark', 'completed', 9, 5, TRUE, 5432),
(12, 1, '2025-03-29', '15:00:00', 'Lindsey Nelson Stadium', 'completed', 11, 2, TRUE, 4267),
(2, 10, '2025-03-30', '18:00:00', 'Baum-Walker Stadium', 'completed', 6, 4, TRUE, 10891),
(7, 3, '2025-04-05', '19:00:00', 'Alex Box Stadium', 'completed', 8, 3, TRUE, 10326),
(15, 16, '2025-04-06', '14:00:00', 'UFCU Disch-Falk Field', 'completed', 9, 6, TRUE, 8112),
(4, 14, '2025-04-12', '18:30:00', 'Florida Ballpark', 'completed', 7, 5, TRUE, 5289),
(12, 5, '2025-04-13', '12:00:00', 'Lindsey Nelson Stadium', 'completed', 10, 4, TRUE, 4356),
(13, 7, '2025-04-19', '19:00:00', 'Blue Bell Park', 'completed', 5, 7, TRUE, 6234),
(2, 1, '2025-04-20', '15:00:00', 'Baum-Walker Stadium', 'completed', 8, 5, TRUE, 11267),
(15, 10, '2025-04-26', '18:00:00', 'UFCU Disch-Falk Field', 'completed', 6, 3, TRUE, 7945),
(4, 12, '2025-04-27', '13:00:00', 'Florida Ballpark', 'completed', 4, 7, TRUE, 5612),
(16, 13, '2025-05-03', '15:00:00', 'L. Dale Mitchell Park', 'completed', 5, 6, TRUE, 3456),

-- Scheduled Games
(7, 2, '2025-05-09', '19:00:00', 'Alex Box Stadium', 'scheduled', NULL, NULL, TRUE, NULL),
(12, 4, '2025-05-10', '18:30:00', 'Lindsey Nelson Stadium', 'scheduled', NULL, NULL, TRUE, NULL),
(13, 15, '2025-05-10', '19:00:00', 'Blue Bell Park', 'scheduled', NULL, NULL, TRUE, NULL),
(14, 2, '2025-05-16', '19:00:00', 'Hawkins Field', 'scheduled', NULL, NULL, TRUE, NULL),
(5, 11, '2025-05-17', '14:00:00', 'Foley Field', 'scheduled', NULL, NULL, TRUE, NULL),
(10, 8, '2025-05-17', '15:00:00', 'Swayze Field', 'scheduled', NULL, NULL, TRUE, NULL),
(1, 3, '2025-05-17', '18:00:00', 'Sewell-Thomas Stadium', 'scheduled', NULL, NULL, TRUE, NULL),
(16, 15, '2025-05-23', '19:00:00', 'L. Dale Mitchell Park', 'scheduled', NULL, NULL, TRUE, NULL),
(7, 4, '2025-05-24', '13:00:00', 'Alex Box Stadium', 'scheduled', NULL, NULL, TRUE, NULL),
(12, 13, '2025-05-24', '18:00:00', 'Lindsey Nelson Stadium', 'scheduled', NULL, NULL, TRUE, NULL);

-- ============================================================
-- Insert Standings (All 16 teams based on their records)
-- ============================================================

INSERT INTO standings (team_id, standing_rank, conference_win_pct, overall_win_pct, home_record, away_record, streak, last_10, games_back, runs_per_game, runs_allowed_per_game) VALUES
(12, 1, 0.8889, 0.8500, '18-2', '12-3', 'W7', '9-1', 0.0, 7.80, 3.90),
(2, 2, 0.8333, 0.8000, '16-2', '13-5', 'W5', '8-2', 1.5, 7.45, 4.18),
(4, 3, 0.7778, 0.7750, '17-3', '11-5', 'W4', '7-3', 2.5, 6.90, 4.45),
(15, 4, 0.7222, 0.7750, '15-4', '13-4', 'W3', '7-3', 3.5, 7.23, 4.30),
(7, 5, 0.7222, 0.7500, '16-3', '11-6', 'W2', '6-4', 3.5, 7.18, 4.63),
(13, 6, 0.6667, 0.7250, '14-5', '12-5', 'L1', '6-4', 5.0, 6.68, 4.73),
(1, 7, 0.6667, 0.7000, '15-6', '10-5', 'W1', '5-5', 5.0, 6.13, 4.95),
(10, 8, 0.6111, 0.6500, '13-7', '10-6', 'W2', '6-4', 6.5, 6.30, 5.03),
(5, 9, 0.6111, 0.6750, '14-6', '10-6', 'L2', '5-5', 6.5, 5.95, 4.80),
(16, 10, 0.6111, 0.6750, '12-5', '12-7', 'W1', '6-4', 6.5, 6.03, 4.88),
(3, 11, 0.5556, 0.6250, '13-7', '9-7', 'L1', '4-6', 8.0, 5.53, 5.13),
(14, 12, 0.5556, 0.6500, '14-6', '9-7', 'W1', '5-5', 8.0, 5.83, 5.08),
(8, 13, 0.5000, 0.6000, '12-8', '9-7', 'L2', '4-6', 9.5, 5.38, 5.25),
(11, 14, 0.5000, 0.5750, '11-8', '9-8', 'L1', '4-6', 9.5, 5.20, 5.43),
(6, 15, 0.4444, 0.5500, '11-9', '8-8', 'L3', '3-7', 11.5, 5.03, 5.58),
(9, 16, 0.3889, 0.5000, '10-10', '7-9', 'L4', '2-8', 13.0, 4.73, 6.03);
