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
import TablePagination from '@mui/material/TablePagination';

function Feature3() {
  const [country, setCountry] = useState("")
  const [topScorers, setTopScorers] = useState([])
  const [rowsPerPage, setRowsPerPage] = useState(10)
  const [count, setCount] = useState()
  const [page, setPage] = useState(0)

  useEffect(() => {
    fetch(`/api/v1/players/topScorer/${country}/?limit=${rowsPerPage}&page=${page + 1}`)
        .then(respose => respose.json())
        .then(json => {
          setTopScorers(json.data)
          setCount(json.total_count)
        })
  }, [country, rowsPerPage, page])

  return (
    <>
      <Typography sx={{m: 2}} variant="h5"><b>Feature 3:</b> Top Goal Scorers</Typography>
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
