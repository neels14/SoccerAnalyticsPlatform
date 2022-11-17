import { Typography } from "@mui/material";
import { useState, useEffect } from "react";
import CountrySelect from "./shared/CountrySelect";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import Paper from '@mui/material/Paper';
import {StyledTableCell, StyledTableRow} from './shared/StyledTable'
import Box from '@mui/material/Box';

function Feature3() {
  const [country, setCountry] = useState("")
  const [topScorers, setTopScorers] = useState([])

  useEffect(() => {
    fetch(`/api/v1/players/topScorer/${country}`)
        .then(respose => respose.json())
        .then(json => setTopScorers(json))
  }, [country])

  return (
    <>
      <Typography sx={{m: 2}} variant="h5"><b>Feature 3:</b> Top Goal Scorers</Typography>
      <Box margin="20px">
        <CountrySelect value={country} onChange={event => setCountry(event.target.value)}/>
      </Box>
      <TableContainer component={Paper}>
          <Table sx={{ minWidth: 650 }}>
              <TableHead>
                  <StyledTableRow>
                      <StyledTableCell>Country</StyledTableCell>
                      <StyledTableCell align="right">Top Scorer</StyledTableCell>
                      <StyledTableCell align="right">Total Goals</StyledTableCell>
                  </StyledTableRow>
              </TableHead>
                <TableBody>
                  {topScorers.map((scorer, index) => (
                      <StyledTableRow key={index}>
                          <StyledTableCell>{scorer.countryName}</StyledTableCell>
                          <StyledTableCell align="right">{scorer.topScorer}</StyledTableCell>
                          <StyledTableCell align="right">{scorer.goals}</StyledTableCell>
                      </StyledTableRow>
                  ))}
                </TableBody>
          </Table>
      </TableContainer>
    </>
  );
}

export default Feature3;
