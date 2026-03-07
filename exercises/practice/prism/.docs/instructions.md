# Instructions

Before activating the laser array, you must predict the exact order in which crystals will be hit, identified by their sample IDs.

## Example Test Case

Consider this crystal array configuration:

```json
{
  "start": { "x": 0, "y": 0, "angle": 0 },
  "prisms": [
    { "id": 3, "x": 30, "y": 10, "angle": 45 },
    { "id": 1, "x": 10, "y": 10, "angle": -90 },
    { "id": 2, "x": 10, "y": 0, "angle": 90 },
    { "id": 4, "x": 20, "y": 0, "angle": 0 }
  ]
}
```

## What's Happening

The laser starts at the origin `(0, 0)` and fires horizontally to the right at angle 0°.
Here's the step-by-step beam path:

**Step 1**: The beam travels along the x-axis (y = 0) and first encounters **Crystal #2** at position `(10, 0)`.
This crystal has a refraction angle of 90°, which means it bends the beam perpendicular to its current path.
The beam, originally traveling at 0°, is now redirected to 90° (straight up).

**Step 2**: The beam now travels vertically upward from position `(10, 0)` and strikes **Crystal #1** at position `(10, 10)`.
This crystal has a refraction angle of -90°, bending the beam by -90° relative to its current direction.
The beam was traveling at 90°, so after refraction it's now at 0° (90° + (-90°) = 0°), traveling horizontally to the right again.

**Step 3**: From position `(10, 10)`, the beam travels horizontally and encounters **Crystal #3** at position `(30, 10)`.
This crystal refracts the beam by 45°, changing its direction to 45°.
The beam continues into empty space beyond the array.
