import { useEffect, useState } from 'react';
import Box from '@mui/material/Box';
import CountrySelect from './shared/CountrySelect';
import ToggleButton from '@mui/material/ToggleButton';
import ToggleButtonGroup from '@mui/material/ToggleButtonGroup';
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import Paper from '@mui/material/Paper';
import {StyledTableCell, StyledTableRow} from './shared/StyledTable'
import { Typography } from '@mui/material';

function Feature1() {
  const [firstOrPodium, setFirstOrPodium] = useState('first');
  const [country, setCountry] = useState("")
  const [totalWins, setTotalWins] = useState([])

  useEffect(() => {
    fetch(`/api/v1/worldCup/${firstOrPodium}/${country}`)
        .then(respose => respose.json())
        .then(json => setTotalWins(json))
  }, [firstOrPodium, country])

  return (
    <>
      <Typography sx={{m: 2}} variant="h5"><b>Feature 1:</b> World Cup Winners and Podiums</Typography>
      <Box display="flex" alignItems="center" margin="20px">
        <FirstPodiumToggle sx={{mr: 5}} value={firstOrPodium} onChange={(event, newFirstOrPodium) => setFirstOrPodium(newFirstOrPodium)}/>
        <CountrySelect value={country} onChange={event => setCountry(event.target.value)}/>
      </Box>
      {totalWins.length ?
        <TableContainer component={Paper}>
            <Table sx={{ minWidth: 650 }}>
                <TableHead>
                    <StyledTableRow>
                        <StyledTableCell>Country</StyledTableCell>
                        <StyledTableCell align="right">Total Wins ({firstOrPodium === "first" ? "First Place" : "Podium"})</StyledTableCell>
                    </StyledTableRow>
                </TableHead>
                    <TableBody>
                      {totalWins.map((country) => (
                          <StyledTableRow key={country.country_name}>
                              <StyledTableCell>{country.country_name}</StyledTableCell>
                              <StyledTableCell align="right">{country.total_win}</StyledTableCell>
                          </StyledTableRow>
                      ))}
                    </TableBody>
            </Table>
        </TableContainer>
        : <Typography sx={{m: 3}} variant="subtitle1">Looks like {country} has never won the World Cup :(</Typography>
      }
    </>
  )
}

function FirstPodiumToggle({...props}) {

  return (
    <ToggleButtonGroup
      color="primary"
      exclusive
      aria-label="Platform"
      {...props}
    >
      <ToggleButton value="first">First Place</ToggleButton>
      <ToggleButton value="podium">Podium</ToggleButton>
    </ToggleButtonGroup>
  );
}

export default Feature1;