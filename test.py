import pexpect
import pexpect.popen_spawn
import sys
import time
import os
 
# Use go run to execute the Go CLI
cli_command = 'go run main.go install'
 
# Check if we're in the correct directory and main.go exists
if not os.path.exists('main.go'):
    print("Error: main.go not found in current directory")
    print("Please run this script from the go-cli directory")
    sys.exit(1)
 
try:
    # Set environment variable to indicate automation
    env = os.environ.copy()
    env['CLI_AUTOMATION'] = 'true'
    
    # Start the Go CLI process using PopenSpawn instead of spawn
    child = pexpect.popen_spawn.PopenSpawn(cli_command, timeout=30, encoding='utf-8', env=env)
    
    # Enable logging of the interaction
    child.logfile = sys.stdout
 
    # Wait for the first prompt
    child.expect('Do you want to install other dependencies.*\(y/N\)')
    # child.send('\b')
    child.sendline('')
    time.sleep(1)
    
    # Wait for the language selection menu
    child.expect('Which language do you want to use.*')
    time.sleep(1)
    
    # Send down arrow key to select JavaScript (second option)
    child.send('\x1b[B')  # Down arrow escape sequence
    time.sleep(0.5)
    
    # Press Enter to confirm selection
    child.sendline('')
    time.sleep(1)
 
    # Wait for username prompt
    child.expect('Enter your username.*')
    time.sleep(0.5)
    child.sendline('Pasindu')
 
    time.sleep(1)
 
    child.expect('Enter your admin name.*')
    child.sendline('Admin')
    time.sleep(1)
 
    # Wait for completion
    child.expect('Installation completed successfully!')
    
    print("\nAutomation completed successfully!")
    
except pexpect.exceptions.EOF as e:
    print(f"Process ended unexpectedly: {e}")
    if 'child' in locals():
        print(f"Command output: {child.before}")
    sys.exit(1)
except pexpect.exceptions.TIMEOUT as e:
    print(f"Process timed out: {e}")
    if 'child' in locals():
        print(f"Command output: {child.before}")
    sys.exit(1)
except Exception as e:
    print(f"Error: {e}")
    if 'child' in locals():
        print(f"Command output: {child.before}")
    sys.exit(1)
 
 
# import pexpect
# import sys
# import time
# import os
 
# # Use go run to execute the Go CLI
# cli_command = 'go run main.go install'
 
# # Check if we're in the correct directory and main.go exists
# if not os.path.exists('main.go'):
#     print("Error: main.go not found in current directory")
#     print("Please run this script from the go-cli directory")
#     sys.exit(1)
 
# try:
#     # Set environment variable to indicate automation
#     env = os.environ.copy()
#     env['CLI_AUTOMATION'] = 'true'
    
#     # Start the Go CLI process using go run
#     child = pexpect.spawn(cli_command, timeout=30, encoding='utf-8', env=env)
    
#     # Enable logging of the interaction
#     child.logfile = sys.stdout
 
#     # Wait for the first prompt
#     child.expect('Do you want to install other dependencies.*\(y/N\)')
#     child.send('\b')
#     child.sendline('y')
#     time.sleep(1)
    
#     # Wait for the language selection menu
#     child.expect('Which language do you want to use.*')
#     time.sleep(1)
    
#     # Send down arrow key to select JavaScript (second option)
#     child.send('\x1b[B')  # Down arrow escape sequence
#     time.sleep(0.5)
    
#     # Press Enter to confirm selection
#     child.sendline('')
#     time.sleep(1)
 
#     # Wait for username prompt
#     child.expect('Enter your username.*')
#     time.sleep(0.5)
#     child.sendline('Pasindu')
 
#     time.sleep(1)
 
#     child.expect('Enter your admin name.*')
#     child.sendline('Admin')
#     child.sendline('')
#     time.sleep(1)
 
#     # Wait for the password prompt
 
#     # # Wait for completion
#     child.expect('Installation completed successfully!')
    
#     print("\nAutomation completed successfully!")
    
# except pexpect.exceptions.EOF as e:
#     print(f"Process ended unexpectedly: {e}")
#     if 'child' in locals():
#         print(f"Command output: {child.before}")
#     sys.exit(1)
# except pexpect.exceptions.TIMEOUT as e:
#     print(f"Process timed out: {e}")
#     if 'child' in locals():
#         print(f"Command output: {child.before}")
#     sys.exit(1)
# except Exception as e:
#     print(f"Error: {e}")
#     if 'child' in locals():
#         print(f"Command output: {child.before}")
#     sys.exit(1)
 