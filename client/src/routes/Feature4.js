import { useEffect, useState } from "react";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import Paper from '@mui/material/Paper';
import {StyledTableCell, StyledTableRow} from './shared/StyledTable'
import { Typography } from "@mui/material";
import Box from "@mui/material/Box";
import CountrySelect from "./shared/CountrySelect";

function Feature4() {
  const [country, setCountry] = useState("")
  const [avgGoals, setAvgGoals] = useState([])

  useEffect(() => {
    fetch(`/api/v1/country/averageGoal/${country}`)
            .then(respose => respose.json())
            .then(json => {if (country) {setAvgGoals([json])} else {setAvgGoals(json)}})
  }, [country])

  return (
    <>
      <Typography sx={{m: 2}} variant="h5"><b>Feature 4:</b> Average Goals Scored Per Match</Typography>
      <Box margin="20px">
        <CountrySelect value={country} onChange={event => setCountry(event.target.value)}/>
      </Box>
      <TableContainer component={Paper}>
        <Table sx={{ minWidth: 650 }}>
            <TableHead>
                <StyledTableRow>
                    <StyledTableCell>Country</StyledTableCell>
                    <StyledTableCell align="right">Average Goals Per Match</StyledTableCell>
                </StyledTableRow>
            </TableHead>
            <TableBody>
                {avgGoals.map((country) => (
                    <StyledTableRow key={country.countryName}>
                        <StyledTableCell>{country.countryName}</StyledTableCell>
                        <StyledTableCell align="right">{country.averageGoal.toFixed(2)}</StyledTableCell>
                    </StyledTableRow>
                ))}
            </TableBody>
        </Table>
      </TableContainer>
    </>
  );
}

export default Feature4;
