const fs = require("fs");

// Function to parse time string (00:00:16,112) to milliseconds
function timeToMs(timeStr) {
  const [time, ms] = timeStr.split(",");
  const [hours, minutes, seconds] = time.split(":").map(Number);
  return hours * 3600000 + minutes * 60000 + seconds * 1000 + Number(ms);
}

// Function to convert milliseconds back to time string
function msToTime(ms) {
  const hours = Math.floor(ms / 3600000);
  const minutes = Math.floor((ms % 3600000) / 60000);
  const seconds = Math.floor((ms % 60000) / 1000);
  const milliseconds = ms % 1000;

  return `${String(hours).padStart(2, "0")}:${String(minutes).padStart(2, "0")}:${String(seconds).padStart(2, "0")},${String(milliseconds).padStart(3, "0")}`;
}

// Function to adjust subtitle timing
function adjustSubtitleTiming(content, delaySeconds) {
  const delayMs = delaySeconds * 1000;
  const lines = content.split("\n");
  const result = [];

  for (let line of lines) {
    // Check if line contains timestamp (format: 00:00:16,112 --> 00:00:21,576)
    if (line.includes("-->")) {
      const [startTime, endTime] = line.split("-->").map((s) => s.trim());

      // Convert to milliseconds, add delay, convert back
      const newStartMs = timeToMs(startTime) + delayMs;
      const newEndMs = timeToMs(endTime) + delayMs;

      const newStartTime = msToTime(newStartMs);
      const newEndTime = msToTime(newEndMs);

      result.push(`${newStartTime} --> ${newEndTime}`);
    } else {
      result.push(line);
    }
  }

  return result.join("\n");
}

// Main execution
try {
  // Read the input file
  const inputContent = fs.readFileSync("princessmononoke.srt", "utf8");

  // Adjust timing by adding 5 seconds
  const adjustedContent = adjustSubtitleTiming(inputContent, 33);

  // Write to output file
  fs.writeFileSync("princessmononoke_out.srt", adjustedContent, "utf8");

  console.log("✓ Successfully adjusted subtitles by +91.5 seconds");
  console.log("✓ Output written to sub2.txt");
} catch (error) {
  console.error("Error:", error.message);
}
