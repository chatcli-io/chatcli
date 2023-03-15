import setuptools

setuptools.setup(
    name='chatcli',
    version="0.0.1",
    description='A cli wrapper around chatgpt, to get tips and hints from the terminal.',
    url='https://github.com/deeliciouscode/chatcli',
    author='David Schmider',
    author_email='schmiderdavid@gmail.com',
    packages=setuptools.find_packages(),
    install_requires=['openai==0.27.2', 'PyYAML==6.0'],
    scripts=['chatcli.py'],
    entry_points={
        'console_scripts': [
            'chatcli=chatcli:run',
        ],
    },
    classifiers=[
        'Programming Language :: Python :: 3.10.9',
        "Operating System :: Unix",
    ],
    python_requires='>=3.7.1',
)
