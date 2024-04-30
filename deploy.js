const fs = require("fs");
const yaml = require("yaml");
const { exec } = require("child_process");

const inventoryPath = "./inventory.yaml"; // Path to the Ansible inventory file

// Function to read and parse the YAML inventory file
function readInventory(path) {
  try {
    const file = fs.readFileSync(path, "utf8");
    return yaml.parse(file);
  } catch (err) {
    console.error("Failed to read or parse the inventory file:", err);
    process.exit(1);
  }
}

// Function to execute Ansible commands
function runAnsible(hosts, jsonData) {
  const envVar = `JSON_DATA='${JSON.stringify(jsonData)}'`;
  const command = `ansible -i ${hosts.join(",")} -m ping`; // Example Ansible command

  exec(`${envVar} ${command}`, (error, stdout, stderr) => {
    if (error) {
      console.error(`Error: ${error.message}`);
      return;
    }
    if (stderr) {
      console.error(`Stderr: ${stderr}`);
      return;
    }
    console.log(`Stdout: ${stdout}`);
  });
}

// Main function to process the inventory and run commands
function processInventory() {
  const inventory = readInventory(inventoryPath);

  Object.keys(inventory).forEach((group) => {
    if (inventory[group].hosts && inventory[group].hosts.length > 0) {
      console.log(`Running Ansible for group: ${group}`);
      runAnsible(inventory[group].hosts, {
        group: group,
        hosts: inventory[group].hosts,
      });
    }
  });
}

processInventory();
