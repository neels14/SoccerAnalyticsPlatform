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
import TablePagination from '@mui/material/TablePagination';

function Feature6() {
  const [country, setCountry] = useState("")
  const [mostStartedPlayers, setMostStartedPlayers] = useState([])
  const [rowsPerPage, setRowsPerPage] = useState(10)
  const [count, setCount] = useState()
  const [page, setPage] = useState(0)

  useEffect(() => {
      fetch(`/api/v1/players/mostStarted/${country}/?limit=${rowsPerPage}&page=${page + 1}`)
          .then(respose => respose.json())
          .then(json => {setMostStartedPlayers(json.data); setCount(json.total_count)})
  }, [country, rowsPerPage, page])

  return (
    <>
      <Typography sx={{m: 2}} variant="h5"><b>Feature 6:</b> Most Started Players</Typography>
      <Box margin="20px">
        <CountrySelect value={country} onChange={event => {setCountry(event.target.value); setPage(0);}}/>
      </Box>
      <TablePagination
          component="div"
          count={count}
          page={page}
          onPageChange={(_event, newPage) => setPage(newPage)}
          rowsPerPage={rowsPerPage}
          onRowsPerPageChange={event => {setRowsPerPage(parseInt(event.target.value, 10)); setPage(0)}}
      />
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
