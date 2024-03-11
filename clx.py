#!python3

import argparse
import json
import requests
import sys
import traceback
import pprint

BASEURL = "http://127.0.0.1:8000/uhppote/simulator"


def commands():
    return {
        "swipe": swipe,
        "swipe-in": swipe_in,
        "swipe-out": swipe_out,
        "open": open_door,
        "close": close_door,
        "button": press_button,
        "list-controllers": list_controllers,
        "create-controller": create_controller,
        "delete-controller": delete_controller,
    }


def exec(f, args):
    response = f(args)
    if response != None:
        pprint.pprint(response.__dict__, indent=2, width=1, sort_dicts=False)


def swipe(args):
    controller = int(f"{args.controller}")
    door = int(f"{args.door}")
    card = int(f"{args.card}")
    PIN = int(f"{args.PIN}") if args.PIN else None

    url = f"{BASEURL}/{controller}/swipe"

    headers = {
        "accept": "application/json",
    }

    body = {
        "door": door,
        "card": card,
        "direction": 1,
    }

    if PIN:
        body["PIN"] = PIN

    response = requests.post(url, headers=headers, json=body)

    if response.ok:
        print("swipe")
        print(json.dumps(response.json(), indent=4))
    else:
        response.raise_for_status()


def swipe_in(args):
    controller = int(f"{args.controller}")
    door = int(f"{args.door}")
    card = int(f"{args.card}")
    PIN = int(f"{args.PIN}") if args.PIN else None

    url = f"{BASEURL}/{controller}/swipe"

    headers = {
        "accept": "application/json",
    }

    body = {
        "door": door,
        "card": card,
        "direction": 1,
    }

    if PIN:
        body["PIN"] = PIN

    response = requests.post(url, headers=headers, json=body)

    if response.ok:
        print("swipe-in")
        print(json.dumps(response.json(), indent=4))
    else:
        response.raise_for_status()


def swipe_out(args):
    controller = int(f"{args.controller}")
    door = int(f"{args.door}")
    card = int(f"{args.card}")
    PIN = int(f"{args.PIN}") if args.PIN else None

    url = f"{BASEURL}/{controller}/swipe"

    headers = {
        "accept": "application/json",
    }

    body = {
        "door": door,
        "card": card,
        "direction": 2,
    }

    if PIN:
        body["PIN"] = PIN

    response = requests.post(url, headers=headers, json=body)

    if response.ok:
        print("swipe-out")
        print(json.dumps(response.json(), indent=4))
    else:
        response.raise_for_status()


def open_door():
    raise ValueError("** NOT IMPLEMENTED **")


def close_door():
    raise ValueError("** NOT IMPLEMENTED **")


def press_button():
    raise ValueError("** NOT IMPLEMENTED **")


def list_controllers():
    raise ValueError("** NOT IMPLEMENTED **")


def create_controller():
    raise ValueError("** NOT IMPLEMENTED **")


def delete_controller():
    raise ValueError("** NOT IMPLEMENTED **")


def parse_args():
    parser = argparse.ArgumentParser(description="uhppote-simulator CLI")

    parser.add_argument(
        "--debug",
        action=argparse.BooleanOptionalAction,
        default=False,
        help="displays detailed error information",
    )

    subparsers = parser.add_subparsers(title="subcommands", dest="command")

    # ... swipe
    swipe = subparsers.add_parser("swipe", help="swipe card")
    swipe.add_argument(
        "--controller", type=int, help="controller serial number, e.g. 405419896"
    )
    swipe.add_argument("--door", type=int, help="door no. [1..4], e.g. 3")
    swipe.add_argument("--card", type=int, help="card number e.g. 10058400")
    swipe.add_argument("--PIN", nargs="?", type=int, help="(optional) card PIN")

    # ... swipe-in
    swipe_in = subparsers.add_parser("swipe-in", help="swipe card IN")
    swipe_in.add_argument(
        "--controller", type=int, help="controller serial number, e.g. 405419896"
    )
    swipe_in.add_argument("--door", type=int, help="door no. [1..4], e.g. 3")
    swipe_in.add_argument("--card", type=int, help="card number e.g. 10058400")
    swipe_in.add_argument("--PIN", nargs="?", type=int, help="(optional) card PIN")

    # ... swipe-out
    swipe_out = subparsers.add_parser("swipe-out", help="swipe card OUT")
    swipe_out.add_argument(
        "--controller", type=int, help="controller serial number, e.g. 405419896"
    )
    swipe_out.add_argument("--door", type=int, help="door no. [1..4], e.g. 3")
    swipe_out.add_argument("--card", type=int, help="card number e.g. 10058400")
    swipe_out.add_argument("--PIN", nargs="?", type=int, help="(optional) card PIN")

    args = parser.parse_args()

    return args


def main():
    if len(sys.argv) < 2:
        usage()
        return -1

    args = parse_args()
    cmd = args.command
    debug = args.debug

    if cmd in commands():
        try:
            exec(commands()[cmd], args)
        except Exception as x:
            print()
            print(f"*** ERROR  {cmd}: {x}")
            print()
            if debug:
                print(traceback.format_exc())
                sys.exit(1)
    else:
        print()
        print(f"  ERROR: invalid command ({cmd})")
        print()
        usage()


def usage():
    print()
    print("  Usage: python3 main.py <command> <args>")
    print()
    print("  Supported commands:")

    for cmd, _ in commands().items():
        print(f"    {cmd}")

    print()


if __name__ == "__main__":
    main()
