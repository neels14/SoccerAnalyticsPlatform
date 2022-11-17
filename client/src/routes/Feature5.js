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

function Feature5() {
  const [country, setCountry] = useState("")
  const [winRatios, setWinRatios] = useState([])

  useEffect(() => {
      fetch(`/api/v1/country/winRatio/${country}`)
          .then(respose => respose.json())
          .then(json => {if (country) {setWinRatios([json])} else {setWinRatios(json)}})
  }, [country])

  return (
    <>
      <Typography sx={{m: 2}} variant="h5"><b>Feature 5:</b> Win Ratio</Typography>
      <Box margin="20px">
        <CountrySelect value={country} onChange={event => setCountry(event.target.value)}/>
      </Box>
      <TableContainer component={Paper}>
          <Table sx={{ minWidth: 650 }}>
              <TableHead>
                  <StyledTableRow>
                      <StyledTableCell>Country</StyledTableCell>
                      <StyledTableCell align="right">Win Ratio</StyledTableCell>
                  </StyledTableRow>
              </TableHead>
              <TableBody>
                  {winRatios.map((winRatio) => (
                      <StyledTableRow key={winRatio.country}>
                          <StyledTableCell>{winRatio.country}</StyledTableCell>
                          <StyledTableCell align="right">{winRatio.winRatio.toFixed(2)}</StyledTableCell>
                      </StyledTableRow>
                  ))}
              </TableBody>
          </Table>
      </TableContainer>
    </>
  );
}

export default Feature5;
