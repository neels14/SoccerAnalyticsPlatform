import { Typography } from "@mui/material";
import { useEffect, useState } from "react";

function Feature2() {
  const [mostPopularYear, setMostPopularYear] = useState()
  const [attendance, setAttendance] = useState("")

  useEffect(() => {
    fetch("/api/v1/worldCup/popular")
        .then(respose => respose.json())
        .then(json => {
          setMostPopularYear(json[0].year)
          setAttendance(json[0].attendance)
        })
  }, [])

  return (
    <>
      <Typography sx={{m: 2}} variant="h5"><b>Feature 2:</b> Most Popular World Cup</Typography>
      <Typography sx={{m: 2}} variant="subtitle1">The most popular World Cup was in <b>{mostPopularYear}</b> with an attendance of <b>{attendance.toLocaleString()}</b></Typography>
    </>
  );
}

export default Feature2;
