import { useEffect, useState } from "react";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import Paper from '@mui/material/Paper';
import {StyledTableCell, StyledTableRow} from './shared/StyledTable'
import TablePagination from '@mui/material/TablePagination';

function Players() {
    const [players, setPlayers] = useState([])
    const [rowsPerPage, setRowsPerPage] = useState(10)
    const [count, setCount] = useState()
    const [page, setPage] = useState(0)

    useEffect(() => {
        fetch(`/api/v1/players/?limit=${rowsPerPage}&page=${page + 1}`)
            .then(respose => respose.json())
            .then(json => {
                setPlayers(json.data)
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
                            <StyledTableCell>Name</StyledTableCell>
                            <StyledTableCell align="right">Country</StyledTableCell>
                            <StyledTableCell align="right">Shirt Number</StyledTableCell>
                        </StyledTableRow>
                    </TableHead>
                    <TableBody>
                        {players.map((player, index) => (
                            <StyledTableRow key={index}>
                                <StyledTableCell>{player.name}</StyledTableCell>
                                <StyledTableCell align="right">{player.country}</StyledTableCell>
                                <StyledTableCell align="right">{player.shirtNumber}</StyledTableCell>
                            </StyledTableRow>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </>
    );
  }
  
export default Players;
