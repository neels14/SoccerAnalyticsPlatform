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

function Feature6() {
  const [country, setCountry] = useState("")
  const [mostStartedPlayers, setMostStartedPlayers] = useState([])

  useEffect(() => {
      fetch(`/api/v1/players/mostStarted/${country}`)
          .then(respose => respose.json())
          .then(json => setMostStartedPlayers(json.data))
  }, [country])

  return (
    <>
      <Typography sx={{m: 2}} variant="h5"><b>Feature 6:</b> Most Started Players</Typography>
      <Box margin="20px">
        <CountrySelect value={country} onChange={event => setCountry(event.target.value)}/>
      </Box>
      <TableContainer component={Paper}>
          <Table sx={{ minWidth: 650 }}>
              <TableHead>
                  <StyledTableRow>
                      <StyledTableCell>Player</StyledTableCell>
                      <StyledTableCell align="right">Country</StyledTableCell>
                      <StyledTableCell align="right">Shirt Number</StyledTableCell>
                      <StyledTableCell align="right">Start Ratio</StyledTableCell>
                  </StyledTableRow>
              </TableHead>
              <TableBody>
                  {mostStartedPlayers.map((mostStartedPlayer, index) => (
                      <StyledTableRow key={index}>
                          <StyledTableCell>{mostStartedPlayer.playerName}</StyledTableCell>
                          <StyledTableCell align="right">{mostStartedPlayer.country}</StyledTableCell>
                          <StyledTableCell align="right">{mostStartedPlayer.shirtNumber}</StyledTableCell>
                          <StyledTableCell align="right">{mostStartedPlayer.startRatio.toFixed(2)}</StyledTableCell>
                      </StyledTableRow>
                  ))}
              </TableBody>
          </Table>
      </TableContainer>
    </>
  );
}

export default Feature6;
