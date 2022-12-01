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

function Feature4() {
  const [country, setCountry] = useState("")
  const [avgGoals, setAvgGoals] = useState([])
  const [rowsPerPage, setRowsPerPage] = useState(10)
  const [count, setCount] = useState()
  const [page, setPage] = useState(0)

  useEffect(() => {
    fetch(`/api/v1/country/averageGoal/${country}/?limit=${rowsPerPage}&page=${page + 1}`)
            .then(respose => respose.json())
            .then(json => {
              if (country) {
                setAvgGoals([json.data])
              } else {
                setAvgGoals(json.data)
              }
              setCount(json.total_count)
            })
  }, [country, rowsPerPage, page])

  return (
    <>
      <Typography sx={{m: 2}} variant="h5"><b>Feature 4:</b> Average Goals Scored Per Match</Typography>
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
