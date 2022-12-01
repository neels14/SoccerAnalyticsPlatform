import { useState, useEffect } from "react";
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select from '@mui/material/Select';

function CountrySelect({...props}) {
    const [countries, setCountries] = useState([])

    useEffect(() => {
        fetch("/api/v1/country/?limit=82")
            .then(respose => respose.json())
            .then(json => setCountries(json.data.map(json => json.name)))
    }, [])

    return (
        <FormControl sx={{minWidth: 150}} size="small">
            <InputLabel id="country-select-label">Country</InputLabel>
            <Select
                labelId="country-select-label"
                id="country-select"
                label="Country"
                {...props}
            >
                <MenuItem value=""><em>All</em></MenuItem>
                {countries.map(country => <MenuItem value={country} key={country}>{country}</MenuItem>)}
            </Select>
        </FormControl>
    );
}

export default CountrySelect;
