import os
import json
import uuid
from datetime import datetime

search_dir = '/home/ma-he-sh/Documents/dev/gitignore'
searchstr  = 'gitignore'
savefile   = './data.json'

def get_hash():
    return str( uuid.uuid4() )

def iterate( dir ):  
    jsondata = {
        'stamp': str( datetime.now() ),
        'hash'   : get_hash(),
        'lang' : {}
    }

    for path, subdirs, files in os.walk( dir ):
        for f in files:
            if searchstr in f:
                fpath = os.path.join( path, f )
                langname = os.path.basename( f )
                langcode = os.path.splitext( langname )[0]
                
                _file = open( fpath, 'r' )
                content = _file.read()

                jsondata['lang'][langcode] = {
                    'fname' : langname,
                    'code'  : langcode,
                    'content': content,
                }

    return jsondata

if __name__ == '__main__':
    print('--parsing started--')
    data = iterate( search_dir )
    
    with open( savefile , 'w' ) as json_file:
        json.dump( data, json_file )

    print('--save file--')
