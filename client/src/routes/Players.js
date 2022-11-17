import { useEffect, useState } from "react";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import Paper from '@mui/material/Paper';
import {StyledTableCell, StyledTableRow} from './shared/StyledTable'

function Players() {
    const [players, setPlayers] = useState([])

    useEffect(() => {
        fetch("/api/v1/players/")
            .then(respose => respose.json())
            .then(json => setPlayers(json))
    }, [])

    return (
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
    );
  }
  
export default Players;
