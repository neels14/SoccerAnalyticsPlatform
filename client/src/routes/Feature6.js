import { useEffect, useState } from "react";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import Paper from '@mui/material/Paper';
import {StyledTableCell, StyledTableRow} from './shared/StyledTable'
import { Typography } from "@mui/material";

function Feature6() {
  const [mostStartedPlayers, setMostStartedPlayers] = useState([])

  useEffect(() => {
      fetch("/api/v1/players/mostStarted")
          .then(respose => respose.json())
          .then(json => setMostStartedPlayers(json))
  }, [])

  return (
    <>
      <Typography sx={{m: 2}} variant="h5"><b>Feature 6:</b> Most Started Players</Typography>
      <TableContainer component={Paper}>
          <Table sx={{ minWidth: 650 }}>
              <TableHead>
                  <StyledTableRow>
                      <StyledTableCell>Country</StyledTableCell>
                      <StyledTableCell align="right">Player</StyledTableCell>
                      <StyledTableCell align="right">Shirt Number</StyledTableCell>
                      <StyledTableCell align="right">Start Ratio</StyledTableCell>
                  </StyledTableRow>
              </TableHead>
              <TableBody>
                  {mostStartedPlayers.map((mostStartedPlayer) => (
                      <StyledTableRow key={mostStartedPlayer.country}>
                          <StyledTableCell>{mostStartedPlayer.country}</StyledTableCell>
                          <StyledTableCell align="right">{mostStartedPlayer.playerName}</StyledTableCell>
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
