import { useEffect, useState } from "react";
import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableContainer from '@mui/material/TableContainer';
import TableHead from '@mui/material/TableHead';
import Paper from '@mui/material/Paper';
import {StyledTableCell, StyledTableRow} from './shared/StyledTable'

function Countries() {
    const [countries, setCountries] = useState([])

    useEffect(() => {
        fetch("/api/v1/country/")
            .then(respose => respose.json())
            .then(json => setCountries(json))
    }, [])

    return (
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
    );
  }
  
export default Countries;
