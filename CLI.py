#!python3

import argparse
import sys
import traceback
import pprint

def commands():
    return {
        'swipe': swipe,
        'swipe-in': swipe_in,
        'swipe-out': swipe_out,
        'open': open_door,
        'close': close_door,
        'button': press_button,
        'list-controllers': list_controllers,
        'create-controller': create_controller,
        'delete-controller': delete_controller,
    }

def exec(f, rest):
    response = f(rest)
    if response != None:
        pprint.pprint(response.__dict__, indent=2, width=1, sort_dicts=False)


def swipe(rest):
    parser = argparse.ArgumentParser(prog='swipe', description='card swipe emulation')
    parser.add_argument('--controller', type=int, help='controller serial number, e.g. 405419896')
    parser.add_argument('--door',       type=int, help='door no. [1..4], e.g. 3')
    parser.add_argument('--card',       type=int, help='card number e.g. 10058400')

    args = parser.parse_args(rest)
    print('>>>>>> ',args.controller)
    print('>>>>>> ',args.door)
    print('>>>>>> ',args.card)

    raise ValueError('** NOT IMPLEMENTED **')

def swipe_in():
    raise ValueError('** NOT IMPLEMENTED **')

def swipe_out():
    raise ValueError('** NOT IMPLEMENTED **')

def open_door():
    raise ValueError('** NOT IMPLEMENTED **')

def close_door():
    raise ValueError('** NOT IMPLEMENTED **')

def press_button():
    raise ValueError('** NOT IMPLEMENTED **')

def list_controllers():
    raise ValueError('** NOT IMPLEMENTED **')

def create_controller():
    raise ValueError('** NOT IMPLEMENTED **')

def delete_controller():
    raise ValueError('** NOT IMPLEMENTED **')

def main():
    if len(sys.argv) < 2:
        usage()
        return -1

    parser = argparse.ArgumentParser(description='uhppote-simulator CLI')

    parser.add_argument('command', 
                        choices=[cmd for cmd in commands().keys()], 
                        help='simulator REST command')
    
    parser.add_argument('--debug',
                        action=argparse.BooleanOptionalAction,
                        default=False,
                        help='displays detailed error information')

    # # ... swipe
    # subparsers = parser.add_subparsers(title='subcommands', dest='command')
    # swipe_parser = subparsers.add_parser('swipe',help='swipe card')
    # swipe_parser.add_argument('--controller', type=int, help='controller serial number, e.g. 405419896')
    # swipe_parser.add_argument('--door',       type=int, help='door no. [1..4], e.g. 3')
    # swipe_parser.add_argument('--card',       type=int, help='card number e.g. 10058400')

    args, rest = parser.parse_known_args()
    cmd = args.command
    debug = args.debug

    if cmd in commands():
        try:
            exec(commands()[cmd](rest))
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
