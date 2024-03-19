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
        'swipe': swipe,
        'swipe-in': swipe_in,
        'swipe-out': swipe_out,
        'passcode': passcode,
        'open': open_door,
        'close': close_door,
        'button': press_button,
        'list-controllers': list_controllers,
        'create-controller': create_controller,
        'delete-controller': delete_controller,
    }


def exec(f, args):
    response = f(args)
    if response != None:
        pprint.pprint(response.__dict__, indent=2, width=1, sort_dicts=False)


def swipe(args):
    controller = args.controller
    door = args.door
    card = args.card
    PIN = args.PIN

    url = f'{BASEURL}/{controller}/swipe'

    headers = {
        'accept': 'application/json',
    }

    body = {
        'door': door,
        'card': card,
        'direction': 1,
    }

    if PIN:
        body['PIN'] = PIN

    response = requests.post(url, headers=headers, json=body)

    if response.ok:
        print('swipe')
        print(json.dumps(response.json(), indent=4))
    else:
        response.raise_for_status()


def swipe_in(args):
    controller = args.controller
    door = args.door
    card = args.card
    PIN = args.PIN

    url = f'{BASEURL}/{controller}/swipe'

    headers = {
        'accept': 'application/json',
    }

    body = {
        'door': door,
        'card': card,
        'direction': 1,
    }

    if PIN:
        body['PIN'] = PIN

    response = requests.post(url, headers=headers, json=body)

    if response.ok:
        print('swipe-in')
        print(json.dumps(response.json(), indent=4))
    else:
        response.raise_for_status()


def swipe_out(args):
    controller = args.controller
    door = args.door
    card = args.card
    PIN = args.PIN

    url = f'{BASEURL}/{controller}/swipe'

    headers = {
        'accept': 'application/json',
    }

    body = {
        'door': door,
        'card': card,
        'direction': 2,
    }

    if PIN:
        body['PIN'] = PIN

    response = requests.post(url, headers=headers, json=body)

    if response.ok:
        print('swipe-out')
        print(json.dumps(response.json(), indent=4))
    else:
        response.raise_for_status()


def passcode(args):
    controller = args.controller
    door = args.door
    code = args.code

    url = f'{BASEURL}/{controller}/code'

    headers = {
        'accept': 'application/json',
    }

    body = {
        'door': door,
        'passcode': code,
    }

    response = requests.post(url, headers=headers, json=body)

    if response.ok:
        print('passcode')
        print(json.dumps(response.json(), indent=4))
    else:
        response.raise_for_status()


def open_door(args):
    controller = args.controller
    door = args.door

    url = f'{BASEURL}/{controller}/door/{door}'

    headers = {
        'accept': 'application/json',
    }

    body = {
        'action': 'open',
    }

    response = requests.post(url, headers=headers, json=body)

    if response.ok:
        print('open-door')
        print(json.dumps(response.json(), indent=4))
    else:
        response.raise_for_status()


def close_door(args):
    controller = args.controller
    door = args.door

    url = f'{BASEURL}/{controller}/door/{door}'

    headers = {
        'accept': 'application/json',
    }

    body = {
        'action': 'close',
    }

    response = requests.post(url, headers=headers, json=body)

    if response.ok:
        print('close-door')
        print(json.dumps(response.json(), indent=4))
    else:
        response.raise_for_status()


def press_button(args):
    controller = args.controller
    door = args.door
    duration = args.duration

    url = f'{BASEURL}/{controller}/door/{door}'

    headers = {
        'accept': 'application/json',
    }

    body = {
        'action': 'button',
        'duration': duration,
    }

    response = requests.post(url, headers=headers, json=body)

    if response.ok:
        print('press-button')
        print(json.dumps(response.json(), indent=4))
    else:
        response.raise_for_status()


def list_controllers(args):
    url = f'{BASEURL}'

    headers = {
        'accept': 'application/json',
    }

    response = requests.get(url, headers=headers)

    if response.ok:
        print('list-controllers')
        print(json.dumps(response.json(), indent=4))
    else:
        response.raise_for_status()


def create_controller(args):
    controller = args.controller
    device_type = args.type
    compressed = args.compressed if args.compressed else False

    url = f'{BASEURL}'

    headers = {
        'accept': 'application/json',
    }

    body = {
        'device-id': controller,
        'device-type': device_type,
        'compressed': compressed,
    }

    response = requests.post(url, headers=headers, json=body)

    if response.ok:
        print('create-controller: created')
    else:
        response.raise_for_status()


def delete_controller(args):
    controller = args.controller

    url = f'{BASEURL}/{controller}'

    headers = {
        'accept': 'application/json',
    }

    response = requests.delete(url, headers=headers)

    if response.ok:
        print('delete-controller: deleted')
    else:
        response.raise_for_status()


def parse_args():
    parser = argparse.ArgumentParser(description='uhppote-simulator CLI')

    parser.add_argument(
        '--debug',
        action=argparse.BooleanOptionalAction,
        default=False,
        help='displays detailed error information',
    )

    subparsers = parser.add_subparsers(title='subcommands', dest='command')

    # ... swipe
    swipe = subparsers.add_parser('swipe', help='swipe card')
    swipe.add_argument('--controller', type=int, help='controller serial number, e.g. 405419896')
    swipe.add_argument('--door', type=int, help='door no. [1..4], e.g. 3')
    swipe.add_argument('--card', type=int, help='card number e.g. 10058400')
    swipe.add_argument('--PIN', nargs='?', type=int, help='(optional) card PIN')

    # ... swipe-in
    swipe_in = subparsers.add_parser('swipe-in', help='swipe card IN')
    swipe_in.add_argument('--controller', type=int, help='controller serial number, e.g. 405419896')
    swipe_in.add_argument('--door', type=int, help='door no. [1..4], e.g. 3')
    swipe_in.add_argument('--card', type=int, help='card number e.g. 10058400')
    swipe_in.add_argument('--PIN', nargs='?', type=int, help='(optional) card PIN')

    # ... swipe-out
    swipe_out = subparsers.add_parser('swipe-out', help='swipe card OUT')
    swipe_out.add_argument('--controller', type=int, help='controller serial number, e.g. 405419896')
    swipe_out.add_argument('--door', type=int, help='door no. [1..4], e.g. 3')
    swipe_out.add_argument('--card', type=int, help='card number e.g. 10058400')
    swipe_out.add_argument('--PIN', nargs='?', type=int, help='(optional) card PIN')

    # ... supervisor passcode
    passcode = subparsers.add_parser('passcode', help='supervisor pass code')
    passcode.add_argument('--controller', type=int, help='controller serial number, e.g. 405419896')
    passcode.add_argument('--door', type=int, help='door no. [1..4], e.g. 3')
    passcode.add_argument('--code', type=int, help='pass code e.g. 123456')

    # ... open door
    open = subparsers.add_parser('open', help='opens door (if unlocked)')
    open.add_argument('--controller', type=int, help='controller serial number, e.g. 405419896')
    open.add_argument('--door', type=int, help='door no. [1..4], e.g. 3')

    # ... close door
    close = subparsers.add_parser('close', help='closes door (if open)')
    close.add_argument('--controller', type=int, help='controller serial number, e.g. 405419896')
    close.add_argument('--door', type=int, help='door no. [1..4], e.g. 3')

    # ... press button
    button = subparsers.add_parser('button', help='emulates a door button press')
    button.add_argument('--controller', type=int, help='controller serial number, e.g. 405419896')
    button.add_argument('--door', type=int, help='door no. [1..4], e.g. 3')
    button.add_argument('--duration', type=int, help='duration of button press (seconds)')

    # ... list controllers
    list_controllers = subparsers.add_parser('list-controllers', help='lists the emulated controllers')

    # ... create controller
    create_controller = subparsers.add_parser('create-controller', help='creates a new controller emulation')
    create_controller.add_argument('--controller', type=int, help='controller serial number, e.g. 405419896')
    create_controller.add_argument('--type', type=str, help='controller type, e.g. UT0311-L04')
    create_controller.add_argument('--compressed', nargs='?', default=False, type=bool, help='store compressed')

    # ... delete controller
    delete_controller = subparsers.add_parser('delete-controller', help='deletes an emulated controller')
    delete_controller.add_argument('--controller', type=int, help='controller serial number, e.g. 405419896')

    # ... parse args
    return parser.parse_args()


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
            print(f'*** ERROR  {cmd}: {x}')
            print()
            if debug:
                print(traceback.format_exc())
                sys.exit(1)
    else:
        print()
        print(f'  ERROR: invalid command ({cmd})')
        print()
        usage()


def usage():
    print()
    print('  Usage: python3 main.py <command> <args>')
    print()
    print('  Supported commands:')

    for cmd, _ in commands().items():
        print(f'    {cmd}')

    print()


if __name__ == '__main__':
    main()
