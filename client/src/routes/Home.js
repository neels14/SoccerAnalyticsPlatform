import { Typography } from "@mui/material";
import Box from "@mui/material/Box";

function Home() {
    return (
        <>
            <Box display="flex" justifyContent="center">
                <img src={require("./world_cup_logo.png")} width="50%" alt="FIFA World Cup"></img>
            </Box>
            <Typography sx={{m: 2}} variant="h4">Welcome to the Soccer Analytics Platform. View data from previous World Cups, check your favourite player's stats and more.</Typography>
            <Typography sx={{m: 2}} variant="h5">Developed by Aashrit Luthra, Anupriyam Ranjit, Hamza Saqib and Neel Shah.</Typography>
        </>
    )
}

export default Home;