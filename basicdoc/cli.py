"""This module provides the OpenApi BasicDoc CLI."""
# basicdoc/cli.py

from pathlib import Path
from typing import Optional
import json
from types import SimpleNamespace
from document import *

import typer

from basicdoc import __app_name__, __version__

app = typer.Typer()

def _version_callback(value: bool) -> None:
    if value:
        typer.echo(f"{__app_name__} v{__version__}")
        raise typer.Exit()

@app.callback()
def main(
    version: Optional[bool] = typer.Option(
        None,
        "--version",
        "-v",
        help="Show the application's version and exit.",
        callback=_version_callback,
        is_eager=True,
    )
) -> None:
    return

@app.command()
def convert(path: str = "./", type: str = "basic", format: str = "json"):
    f = open(path + '/info.json')
    data = json.loads(f)
    # Parse JSON into an object with attributes corresponding to dict keys.
    # x = json.loads(data, object_hook=lambda d: SimpleNamespace(**))
    a = BasicInfo(**data)
    print(a)

if __name__ == "__main__":
    app()