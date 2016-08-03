#Introduction to Antennas and Radio Waves

Date: 1-2-2016

## Electromagnetic Waves
- Electricity (blues)
- Magnetism (red)
- Electric field is perpendicular to the magnetic field
- Doesn't require a medium
- Measured as both particle and wave properties
	- Electronic volts
	- Hertz
- Speeds 180,000 miles/second or 3x10^8 meters/second
- Wavelength is lambda (wavelength in meters) = V (velocity)/F (frequency in hertz)
- Lot of IoT/FCC fuckery in VHF (30 to 300MHz) and UHF (300 to 3,000 MHz)

## Radio Wave Propagation
- Ground Waves travel along the surface
- Sky Waves travel from the transmitter to one of the ionospheric layers and is returned to the earth (art vs science) (can be very low power)
- Reflection (same angle it strikes)
- Refraction (bending through)
- Diffraction (bending around)
- Absorption (agitation of object's matter)
- Scatter (multi-directional bounces)
- Man-made and natural objs
- Critical Freq
- Minimum power for transmission (based upon receiver's condition) (transition.fcc.gov/pshs/techtopics/techtopics17.html);
- * HF is preferred since it is a good combination of distance and power required *
- SOMETIMES THE SUN CAN FUCK YOUR SHIT UP (aka sunspots), thus a lot of * important * stuff is done at night
- Can use low power on an in use frequency w/ modulation to mask signal

## Antenna Basics
- Uses a transmission line to transmit/receive/tx&rx data via a carrier that tries to get to "ground"
- Antenna length: constant / frequency (MHz) = Antenna length in feeet and inches (or meters)
- Lower frequency, longer antenna; higher, shorter
- Only good for a certain range
- Consult a Smith Chart
- Standing wave ratio (when properly cut sine wave begins at one end of the antenna and ends at the other)
- Polarization: depends on the orientation of the emission vector for the wave in relationship to the earth
- Polarity matters more at longer distances
- Antenna Gain: Directly related to directivity; Inverse Square Law (Signal strength for a given surface area is inversly proportional to the distance it traveled from source)
- Antenna Loss: Impedance everywhere; Both internal to the system and external to the antenna
- The FCC will fuck you, so do it right
- Interference mitigation via AGC, filters, design, location selection, etc
- Transmission Line: Length (in meters) = (300/ Center Freq) * VF; coax (minimum loss, loss increases w/ freq, shielded) is v popular right now; Power Loss (radiation, heating, reflection)
- Antenna Counterpoise: Prevent ground reflected waves from absorption

## Common Antennas
- Horizontal Antennas: Suggested max resistance is 100ohms; 1/4 Wavelength above the ground; Practical Half-wave antennas; Phase reversal takes place; sloping wire setup is v potent for HF
- Grounded Antennas: 1/4 wave antennas; used w/ highly conductive ground; ground rod uses the earth as the other 1/4 wavelength element
- Ground Plane Antennas: 1 section verticle 1/4 wavelength; N number of spokes 1/4 wavelength (form an aritifical ground, VHF or higher freqs.); Example: disc-cone antenna; CB in trucks use this
- Fractal Antenna: multi-level/ space filling; popular for RFID; uses repeating iterations
- "Cantenna": Directional wave guide antenna; common for WiFi applications

## Misc
- Laws make things not fun so don't get caught
- Deal w/ electricity in a safe manner
- When in doubt ground it

## Links
[AMSAT](http://www.amsat.org/)
[DXZone](http://www.dxzone.com/)
