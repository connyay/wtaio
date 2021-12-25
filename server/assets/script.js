async function getClosestFeature(lat, long) {
  const res = await fetch(`/${lat}/${long}`);
  return await res.text();
}

async function getCoordinates() {
  if (location.hostname === "0.0.0.0") {
    return { latitude: 39.637833, longitude: -106.047516 };
  }
  const { coords } = await new Promise((resolve, reject) => {
    navigator.geolocation.getCurrentPosition(resolve, reject, {
      enableHighAccuracy: true,
      timeout: 5000,
      maximumAge: 0,
    });
  });
  return coords;
}

async function tick() {
  try {
    const coords = await getCoordinates();
    console.log("Your current position is:");
    console.log(`Latitude : ${coords.latitude}`);
    console.log(`Longitude: ${coords.longitude}`);

    const closest = await getClosestFeature(coords.latitude, coords.longitude);
    console.log(`The closest thing is ${closest}`);
    document.getElementById("closest").innerText = closest;
    document.getElementById("coords").innerText = JSON.stringify(
      {
        latitude: coords.latitude,
        longitude: coords.longitude,
      },
      null,
      2
    );
  } finally {
    setTimeout(tick, 1000);
  }
}

tick();
