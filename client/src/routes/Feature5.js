import { useEffect, useState } from "react";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import Paper from '@mui/material/Paper';
import {StyledTableCell, StyledTableRow} from './shared/StyledTable'
import { Typography } from "@mui/material";

function Feature5() {
  const [winRatios, setWinRatios] = useState([])

  useEffect(() => {
      fetch("/api/v1/country/winRatio")
          .then(respose => respose.json())
          .then(json => setWinRatios(json))
  }, [])

  return (
    <>
      <Typography sx={{m: 2}} variant="h5"><b>Feature 5:</b> Win Ratio</Typography>
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
