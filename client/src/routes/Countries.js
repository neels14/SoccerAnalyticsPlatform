import { useEffect, useState } from "react";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import Paper from '@mui/material/Paper';
import {StyledTableCell, StyledTableRow} from './shared/StyledTable'
import TablePagination from '@mui/material/TablePagination';

function Countries() {
    const [countries, setCountries] = useState([])
    const [rowsPerPage, setRowsPerPage] = useState(10)
    const [count, setCount] = useState()
    const [page, setPage] = useState(0)

    useEffect(() => {
        fetch(`/api/v1/country/?limit=${rowsPerPage}&page=${page + 1}`)
            .then(respose => respose.json())
            .then(json => {
                setCountries(json.data)
                setCount(json.total_count)
            })
    }, [rowsPerPage, page])

    return (
        <>
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
                            <StyledTableCell align="right">Team Initials</StyledTableCell>
                        </StyledTableRow>
                    </TableHead>
                    <TableBody>
                        {countries.map((country) => (
                            <StyledTableRow key={country.name}>
                                <StyledTableCell>{country.name}</StyledTableCell>
                                <StyledTableCell align="right">{country.teamInitial}</StyledTableCell>
                            </StyledTableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </>
    );
  }
  
export default Countries;
