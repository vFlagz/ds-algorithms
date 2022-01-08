package main

import "fmt"

func main() {
	entries := getEntry()
	result := findTwoNumber(entries)
	fmt.Println("rewsult is ??A", result)
}

func findTwoNumber(entry []int) int {
	for i := 0; i < len(entry); i++ {
		for j := 0; j < len(entry); j++ {
			result := entry[i] + entry[j]
			if result == 2020 {
				return multiplyTwoNumber(entry[i], entry[j])
			}
		}
	}
	return -1
}

func multiplyTwoNumber(a int, b int) int {
	return a * b
}

func getEntry() []int {
	testEntries := []int{1695, 1157, 1484, 1717, 622, 1513, 1924, 63, 1461, 1971, 1382, 1587, 1913, 1665, 1464, 1914, 1637, 1527, 1424, 1361, 1187, 272, 1909, 1448, 1623, 1164, 1931, 1646, 1096, 1655, 1962, 1961, 1694, 1792, 1989, 1616, 138, 1887, 1357, 1965, 1085, 308, 2007, 1254, 1179, 1124, 1719, 1467, 1928, 1630, 1676, 1359, 1241, 1511, 1413, 1656, 1818, 1919, 1422, 1745, 1208, 1609, 1544, 1775, 1154, 1057, 1440, 1242, 1202, 1266, 1305, 1836, 1760, 1730, 1396, 1315, 1496, 1964, 1300, 1195, 1583, 1607, 1743, 1682, 1453, 1848, 1320, 1601, 954, 1473, 1847, 1486, 1853, 1668, 1342, 1087, 1139, 1349, 1568, 1728, 1420, 1233, 1073, 1376, 1658, 1477, 1871, 1958, 1950, 1503, 1758, 1474, 1203, 1336, 1981, 1309, 1618, 1846, 1974, 1940, 1333, 1119, 1756, 1918, 961, 1307, 1375, 1346, 1611, 1284, 84, 1754, 1608, 2010, 1341, 1136, 1218, 1882, 1911, 1288, 1930, 1749, 1952, 1556, 1757, 1761, 1112, 1963, 1186, 1373, 1622, 1973, 1330, 1508, 1222, 1226, 1389, 1679, 1584, 1237, 1563, 1763, 1998, 1293, 1642, 95, 1661, 1674, 1100, 1262, 1895, 1548, 1400, 1205, 1435, 1156, 1034, 1577, 1701, 1198, 1173, 1500, 1858, 1809, 1780, 1412, 1982, 1070, 1523, 1776, 1598, 1113, 1144, 1777, 1313, 1102, 1999, 1405, 1784, 1196}
	return testEntries
}
