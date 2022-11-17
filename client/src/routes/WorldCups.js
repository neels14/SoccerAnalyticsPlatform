import { useEffect, useState } from "react";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import Paper from '@mui/material/Paper';
import {StyledTableCell, StyledTableRow} from './shared/StyledTable'

function WorldCups() {
    const [worldCups, setWorldCups] = useState([])

    useEffect(() => {
        fetch("/api/v1/worldCup/")
            .then(respose => respose.json())
            .then(json => setWorldCups(json))
    }, [])

    return (
        <TableContainer component={Paper}>
            <Table sx={{ minWidth: 650 }}>
                <TableHead>
                    <StyledTableRow>
                        <StyledTableCell>Year</StyledTableCell>
                        <StyledTableCell align="right">Host Country</StyledTableCell>
                        <StyledTableCell align="right">First Place</StyledTableCell>
                        <StyledTableCell align="right">Second Place</StyledTableCell>
                        <StyledTableCell align="right">Third Place</StyledTableCell>
                    </StyledTableRow>
                </TableHead>
                <TableBody>
                    {worldCups.map((worldCup) => (
                        <StyledTableRow key={worldCup.year}>
                            <StyledTableCell>{worldCup.year}</StyledTableCell>
                            <StyledTableCell align="right">{worldCup.host}</StyledTableCell>
                            <StyledTableCell align="right">{worldCup.first}</StyledTableCell>
                            <StyledTableCell align="right">{worldCup.second}</StyledTableCell>
                            <StyledTableCell align="right">{worldCup.third}</StyledTableCell>
                        </StyledTableRow>
                    ))}
                </TableBody>
            </Table>
        </TableContainer>
    );
  }
  
export default WorldCups;
