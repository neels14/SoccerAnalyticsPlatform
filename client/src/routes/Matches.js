import { useEffect, useState } from "react";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import Paper from '@mui/material/Paper';
import {StyledTableCell, StyledTableRow} from './shared/StyledTable'
import TablePagination from '@mui/material/TablePagination';

function Matches() {
    const [matches, setMatches] = useState([])
    const [rowsPerPage, setRowsPerPage] = useState(10)
    const [count, setCount] = useState()
    const [page, setPage] = useState(0)

    useEffect(() => {
        fetch(`/api/v1/matches/?limit=${rowsPerPage}&page=${page + 1}`)
            .then(respose => respose.json())
            .then(json => {
                setMatches(json.data)
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
                            <StyledTableCell>Year</StyledTableCell>
                            <StyledTableCell align="right">Stage</StyledTableCell>
                            <StyledTableCell align="right">Home Team</StyledTableCell>
                            <StyledTableCell align="right">Away Team</StyledTableCell>
                            <StyledTableCell align="right">City</StyledTableCell>
                            <StyledTableCell align="right">Stadium</StyledTableCell>
                            <StyledTableCell align="right">Date</StyledTableCell>
                            <StyledTableCell align="right">Attendance</StyledTableCell>
                        </StyledTableRow>
                    </TableHead>
                    <TableBody>
                        {matches.map((match, index) => (
                            <StyledTableRow key={index}>
                                <StyledTableCell>{match.year}</StyledTableCell>
                                <StyledTableCell align="right">{match.stage}</StyledTableCell>
                                <StyledTableCell align="right">{match.home}</StyledTableCell>
                                <StyledTableCell align="right">{match.away}</StyledTableCell>
                                <StyledTableCell align="right">{match.city}</StyledTableCell>
                                <StyledTableCell align="right">{match.stadium}</StyledTableCell>
                                <StyledTableCell align="right">{match.date}</StyledTableCell>
                                <StyledTableCell align="right">{match.attendance}</StyledTableCell>
                            </StyledTableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </>
    );
  }
  
export default Matches;
