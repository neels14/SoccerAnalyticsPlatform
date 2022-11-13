-- name: CSV-DROP-WorldCupsCSV
DROP TABLE IF EXISTS WorldCupsCSV;
-- name: CSV-Table-WorldCupsCSV
CREATE TABLE IF NOT EXISTS WorldCupsCSV(
   Year           INTEGER  NOT NULL
  ,Country        VARCHAR(12) NOT NULL
  ,Winner         VARCHAR(10) NOT NULL
  ,RunnersUp      VARCHAR(14) NOT NULL
  ,Third          VARCHAR(11) NOT NULL
  ,Fourth         VARCHAR(14) NOT NULL
  ,GoalsScored    INTEGER  NOT NULL
  ,QualifiedTeams INTEGER  NOT NULL
  ,MatchesPlayed  INTEGER  NOT NULL
  ,Attendance     VARCHAR(9) NOT NULL
);
-- name: CSV-Table-WorldCupsCSV-Insert
INSERT INTO WorldCupsCSV(Year,Country,Winner,RunnersUp,Third,Fourth,GoalsScored,QualifiedTeams,MatchesPlayed,Attendance) VALUES
 (1930,'Uruguay','Uruguay','Argentina','USA','Yugoslavia',70,13,18,'590.549')
,(1934,'Italy','Italy','Czechoslovakia','Germany','Austria',70,16,17,'363.000')
,(1938,'France','Italy','Hungary','Brazil','Sweden',84,15,18,'375.700')
,(1950,'Brazil','Uruguay','Brazil','Sweden','Spain',88,13,22,'1.045.246')
,(1954,'Switzerland','Germany FR','Hungary','Austria','Uruguay',140,16,26,'768.607')
,(1958,'Sweden','Brazil','Sweden','France','Germany FR',126,16,35,'819.810')
,(1962,'Chile','Brazil','Czechoslovakia','Chile','Yugoslavia',89,16,32,'893.172')
,(1966,'England','England','Germany FR','Portugal','Soviet Union',89,16,32,'1.563.135')
,(1970,'Mexico','Brazil','Italy','Germany FR','Uruguay',95,16,32,'1.603.975')
,(1974,'Germany','Germany FR','Netherlands','Poland','Brazil',97,16,38,'1.865.753')
,(1978,'Argentina','Argentina','Netherlands','Brazil','Italy',102,16,38,'1.545.791')
,(1982,'Spain','Italy','Germany FR','Poland','France',146,24,52,'2.109.723')
,(1986,'Mexico','Argentina','Germany FR','France','Belgium',132,24,52,'2.394.031')
,(1990,'Italy','Germany FR','Argentina','Italy','England',115,24,52,'2.516.215')
,(1994,'USA','Brazil','Italy','Sweden','Bulgaria',141,24,52,'3.587.538')
,(1998,'France','France','Brazil','Croatia','Netherlands',171,32,64,'2.785.100')
,(2002,'Japan','Brazil','Germany','Turkey','Korea Republic',161,32,64,'2.705.197')
,(2006,'Germany','Italy','France','Germany','Portugal',147,32,64,'3.359.439')
,(2010,'South Africa','Spain','Netherlands','Germany','Uruguay',145,32,64,'3.178.856')
,(2014,'Brazil','Germany','Argentina','Netherlands','Brazil',171,32,64,'3.386.810');
