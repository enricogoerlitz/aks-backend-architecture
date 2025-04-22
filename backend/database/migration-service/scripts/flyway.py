import os
import subprocess
import argparse


CONFIG_FILE_NAME = "flyway.conf"
BASE_DIR = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
CONFIG_FILE = os.path.join(BASE_DIR, CONFIG_FILE_NAME)


def run_flyway_command(command, config_file):
    """Führt einen Flyway-Befehl aus."""
    flyway_cmd = ["flyway", f"-configFiles={config_file}", command]

    print(f"Ausführen: {' '.join(flyway_cmd)}")
    result = subprocess.run(flyway_cmd, capture_output=True, text=True)

    print(result.stdout)
    if result.stderr:
        print(f"ERROR: {result.stderr}")

    return result.returncode == 0


def main():
    parser = argparse.ArgumentParser(description="Flyway Migrations Manager")
    parser.add_argument("command", choices=["migrate", "clean", "info", "validate", "repair", "baseline"],
                        help="Auszuführender Flyway-Befehl")

    args = parser.parse_args()

    if not os.path.exists(CONFIG_FILE):
        print(f"Konfigurationsdatei '{CONFIG_FILE}' nicht gefunden für Umgebung '{args.system_env}'")
        exit(1)

    success = run_flyway_command(args.command, CONFIG_FILE)

    if success:
        print(f"Flyway {args.command} erfolgreich ausgeführt")
    else:
        print(f"Flyway {args.command} fehlgeschlagen")
        exit(1)


if __name__ == "__main__":
    main()
